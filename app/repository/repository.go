package repository

import (
	"gorm.io/gorm"
	"road_to_mixi/models"
)

func GetUserByID(db *gorm.DB, id int) (models.User, error) {
    var user models.User
    return user, db.First(&user, id).Error
}

func Get_friend_list(db *gorm.DB, id int) ([]models.Friend, error) {
	var friends []models.Friend
	return friends, db.Model(&models.FriendLink{}).
		Select("User2.user_id AS id, User2.name").
		Joins("User2").
		Where("friend_links.user1_id = ?", id).
		Scan(&friends).Error
}

func Get_friend_of_friend_list(db *gorm.DB, id int) ([]models.Friend, error) {
	var friends []models.Friend
	subQuery := db.Model(&models.FriendLink{}).
		Select("user2_id").
		Where("user1_id = ?", id)
	var blockees []int
	var blockers []int
	db.Model(&models.BlockList{}).
		Where("user1_id = ?", id).
		Pluck("user2_id", &blockees)
	db.Model(&models.BlockList{}).
		Where("user2_id = ?", id).
		Pluck("user1_id", &blockers)
	blockedIDs := append(blockees, blockers...)
	if err := db.Model(&models.FriendLink{}).
		Distinct("User2.user_id").
		Select("User2.user_id AS id, User2.name AS name").
		Joins("User2").
		Where("friend_links.user1_id IN (?)", subQuery).
		Where("friend_links.user2_id NOT IN (?)", subQuery).
		Where("friend_links.user2_id NOT IN (?)", subQuery).
		Where("friend_links.user2_id != ?", id).
		Where("friend_links.user2_id NOT IN (?)", blockedIDs).
		Scan(&friends).Error; err != nil {
		return nil, err
	}
	return friends, nil
}

func Get_friend_of_friend_list_paging(db *gorm.DB, id int, page int, limit int) ([]models.Friend, error) {
	var blockees []int
	var blockers []int
	var friends []models.Friend
	subQuery := db.Model(&models.FriendLink{}).
		Select("user2_id").
		Where("user1_id = ?", id)
	db.Model(&models.BlockList{}).
		Where("user1_id = ?", id).
		Pluck("user2_id", &blockees)
	db.Model(&models.BlockList{}).
		Where("user2_id = ?", id).
		Pluck("user1_id", &blockers)
	blockedIDs := append(blockees, blockers...)

	if err := db.Model(&models.FriendLink{}).
		Distinct("User2.user_id").
		Select("User2.user_id AS id, User2.name AS name").
		Joins("User2").
		Where("friend_links.user1_id IN (?)", subQuery).
		Where("friend_links.user2_id != ?", id).
		Where("friend_links.user2_id NOT IN (?)", subQuery).
		Where("friend_links.user2_id NOT IN (?)", blockedIDs).
		Offset((page - 1) * limit).
		Limit(limit).
		Scan(&friends).Error; err != nil {
		return nil, err
	}
	return friends, nil
}
