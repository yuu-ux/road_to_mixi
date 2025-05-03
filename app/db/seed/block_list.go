package seed

import (
	"gorm.io/gorm"
	"log"
	"road_to_mixi/models"
)

func SeedBlockLists(db *gorm.DB) {
	blockList := []models.BlockList{
		{User1ID: 2, User2ID: 1},
		{User1ID: 3, User2ID: 5},
		{User1ID: 4, User2ID: 2},
		{User1ID: 5, User2ID: 7},
		{User1ID: 6, User2ID: 3},
		{User1ID: 7, User2ID: 4},
	}

	if err := db.Create(&blockList).Error; err != nil {
		log.Fatalf("Failed to seed: %v", err)
	}
	log.Println("User seeding completed successfully.")
}
