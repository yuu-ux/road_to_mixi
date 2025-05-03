package seed

import (
	"gorm.io/gorm"
	"log"
	"road_to_mixi/models"
)

func SeedUsers(db *gorm.DB) {
	user := []models.User{
		{UserID: 1, Name: "タカシ"},
		{UserID: 2, Name: "ユキ"},
		{UserID: 3, Name: "ハルカ"},
		{UserID: 4, Name: "ケンジ"},
		{UserID: 5, Name: "アユミ"},
		{UserID: 6, Name: "サトシ"},
		{UserID: 7, Name: "ミサキ"},
		{UserID: 8, Name: "ダイスケ"},
		{UserID: 9, Name: "ケイコ"},
		{UserID: 10, Name: "リュウタ"},
		{UserID: 11, Name: "マユコ"},
		{UserID: 12, Name: "ケンタロウ"},
		{UserID: 13, Name: "サオリ"},
		{UserID: 14, Name: "タクヤ"},
		{UserID: 15, Name: "アケミ"},
		{UserID: 16, Name: "ナオキ"},
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Failed to seed: %v", err)
	}
	log.Println("User seeding completed successfully.")
}
