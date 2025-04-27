package seed

import (
	"log"
	"gorm.io/gorm"
    "road_to_mixi/models"
)

func SeedUsers(db *gorm.DB) {
	user := []models.User{
		{UserID: 1, Name: "Alice"},
		{UserID: 2, Name: "Bob"},
		{UserID: 3, Name: "Charlie"},
		{UserID: 4, Name: "David"},
		{UserID: 5, Name: "Eve"},
		{UserID: 6, Name: "Frank"},
	}

    if err := db.Create(&user).Error; err != nil {
        log.Fatalf("Failed to seed: %v", err)
    }
	log.Println("User seeding completed successfully.")
}
