package seed

import (
	"log"
	"gorm.io/gorm"
    "road_to_mixi/models"
)

func SeedBlockLists(db *gorm.DB) {
	blockList := []models.BlockList{
		{User1ID: 1, User2ID: 3},
		{User1ID: 2, User2ID: 1},
		{User1ID: 3, User2ID: 2},
	}

    db.Exec("TRUNCATE TABLE block_lists")
    if err := db.Create(&blockList).Error; err != nil {
        log.Fatalf("Failed to seed: %v", err)
    }
	log.Println("User seeding completed successfully.")
}
