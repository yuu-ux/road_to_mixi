package seed

import (
	"log"
	"gorm.io/gorm"
    "road_to_mixi/models"
)

func SeedBlockLists(db *gorm.DB) {
	blockList := []models.BlockList{
		{User1ID: 2, User2ID: 1},
	}

    if err := db.Create(&blockList).Error; err != nil {
        log.Fatalf("Failed to seed: %v", err)
    }
	log.Println("User seeding completed successfully.")
}
