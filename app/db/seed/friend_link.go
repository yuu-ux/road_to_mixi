package seed

import (
	"gorm.io/gorm"
	"log"
	"road_to_mixi/models"
)

func SeedFriendLinks(db *gorm.DB) {
	friendLink := []models.FriendLink{
		{User1ID: 1, User2ID: 2},
		{User1ID: 1, User2ID: 6},
		{User1ID: 1, User2ID: 3},
		{User1ID: 1, User2ID: 4},
		{User1ID: 2, User2ID: 3},
		{User1ID: 2, User2ID: 5},
		{User1ID: 2, User2ID: 4},
		{User1ID: 2, User2ID: 8},
		{User1ID: 3, User2ID: 4},
		{User1ID: 3, User2ID: 1},
		{User1ID: 3, User2ID: 6},
		{User1ID: 3, User2ID: 5},
		{User1ID: 4, User2ID: 5},
		{User1ID: 4, User2ID: 2},
		{User1ID: 4, User2ID: 1},
		{User1ID: 5, User2ID: 6},
		{User1ID: 5, User2ID: 2},
		{User1ID: 6, User2ID: 4},
		{User1ID: 6, User2ID: 3},
		{User1ID: 6, User2ID: 7},
	}

	if err := db.Create(&friendLink).Error; err != nil {
		log.Fatalf("Failed to seed: %v", err)
	}
	log.Println("User seeding completed successfully.")
}
