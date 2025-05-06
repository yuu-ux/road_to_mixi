package test

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	database "road_to_mixi/db"
	"road_to_mixi/models"
	"road_to_mixi/repository"
	"testing"
)

func setupTestDB() *gorm.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.New()
	if err != nil {
		panic("Failed to connect to MySQL: " + err.Error())
	}
	db.Exec("DROP TABLE IF EXISTS friend_links")
	db.Exec("DROP TABLE IF EXISTS block_lists")
	db.Exec("DROP TABLE IF EXISTS users")
	db.AutoMigrate(&models.User{}, &models.FriendLink{}, &models.BlockList{})
	database.InitDatabase(db)
	return db
}

func TestGetFriendList(t *testing.T) {
	db := setupTestDB()
	friends, err := repository.Get_friend_list(db, "1")
	if err != nil {
		t.Fatalf("Error fetching friend list: %v", err)
	}
	if len(friends) != 2 {
		t.Errorf("Expected 2 friend, got %d", len(friends))
	}
}

func TestGetFriendOfFriendList(t *testing.T) {
	db := setupTestDB()
	friends, err := repository.GetFriendOfFriendList(db, "1")
	if err != nil {
		t.Fatalf("Error fetching friend of friend list: %v", err)
	}
	if len(friends) != 2 {
		t.Errorf("Expected 2 friend of friend, got %d", len(friends))
	}
}

func TestGetFriendOfFriendListPaging(t *testing.T) {
	db := setupTestDB()
	friends, err := repository.GetFriendOfFriendListPaging(db, "1", 2, 1)
	if err != nil {
		t.Fatalf("Error fetching paginated friend of friend list: %v", err)
	}
	if len(friends) != 1 {
		t.Errorf("Expected 1 friend of friend in page, got %d", len(friends))
	}
}
