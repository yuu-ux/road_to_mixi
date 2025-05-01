package db

import (
	"fmt"
	"os"
	"road_to_mixi/db/seed"
	"road_to_mixi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    database := os.Getenv("DB_DATABASE")
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, password, host, port, database)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

func InitDatabase(db *gorm.DB) error {
    if err := db.AutoMigrate(&models.User{}, &models.FriendLink{}, &models.BlockList{}); err != nil {
        return fmt.Errorf("auto migrate failed: %w", err)
    }
    tables := []string{"users", "friend_links", "block_lists"}
    if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
        return fmt.Errorf("failed to disable foreign key checks: %w", err)
    }
    for _, t := range tables {
        if db.Migrator().HasTable(t) {
            if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", t)).Error; err != nil {
                return fmt.Errorf("failed to truncate table %s: %w", t, err)
            }
        }
    }
    if err := db.Exec("SET FOREIGN_KEY_CHECKS = 1").Error; err != nil {
        return fmt.Errorf("failed to enable foreign key checks: %w", err)
    }
    seed.SeedUsers(db)
    seed.SeedBlockLists(db)
    seed.SeedFriendLinks(db)
	return nil
}

