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

	// 正常ケース
	friends, err := repository.GetFriendList(db, 1)
	if err != nil {
		t.Fatalf("Error fetching friend list: %v", err)
	}
	expectedFriendIDs := []int{2, 3, 4, 6}
	for _, friend := range friends {
		found := false
		for _, expectedID := range expectedFriendIDs {
			if friend.ID == expectedID {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected friend ID: %d, not in expected list", friend.ID)
		}
	}

	// 異常ケース
	friends, err = repository.GetFriendList(db, 999)
	if err != nil {
		t.Fatalf("Error fetching friend list: %v", err)
	}
	if len(friends) != 0 {
		t.Errorf("Expected 0 friend, got %d", len(friends))
	}
}

func TestGetFriendOfFriendList(t *testing.T) {
	db := setupTestDB()

	// 正常ケース
	friends, err := repository.GetFriendOfFriendList(db, 1)
	if err != nil {
		t.Fatalf("Error fetching friend of friend list: %v", err)
	}
	expectedFriendIDs := []int{7}
	for _, friend := range friends {
		found := false
		for _, expectedID := range expectedFriendIDs {
			if friend.ID == expectedID {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected friend ID: %d, not in expected list", friend.ID)
		}
	}

	// 異常ケース
	friends, err = repository.GetFriendOfFriendList(db, 999)
	if err != nil {
		t.Fatalf("Error fetching friend of friend list: %v", err)
	}
	if len(friends) != 0 {
		t.Errorf("Expected 0 friend of friend, got %d", len(friends))
	}
}

func TestGetFriendOfFriendListPaging(t *testing.T) {
	db := setupTestDB()

	// 正常ケース
	friends, err := repository.GetFriendOfFriendListPaging(db, 1, 2, 1)
	if err != nil {
		t.Fatalf("Error fetching paginated friend of friend list: %v", err)
	}
	expectedFriendIDs := []int{7}
	for _, friend := range friends {
		found := false
		for _, expectedID := range expectedFriendIDs {
			if friend.ID == expectedID {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected friend ID: %d, not in expected list", friend.ID)
		}
	}

	friends, err = repository.GetFriendOfFriendList(db, 999)
	if err != nil {
		t.Fatalf("Error fetching friend of friend list: %v", err)
	}
	// 異常ケース
	if len(friends) != 0 {
		t.Errorf("Expected 0 friend of friend, got %d", len(friends))
	}
}
