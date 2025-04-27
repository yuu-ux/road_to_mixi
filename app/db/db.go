package db

import (
	"fmt"
	"os"
	"road_to_mixi/db/seed"

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

func InitDatabase(db *gorm.DB) {
    db.Exec("SET FOREIGN_KEY_CHECKS = 0")
    db.Exec("TRUNCATE TABLE users")
    db.Exec("TRUNCATE TABLE friend_links")
    db.Exec("TRUNCATE TABLE block_lists")
    db.Exec("SET FOREIGN_KEY_CHECKS = 1")
    seed.SeedUsers(db)
    seed.SeedBlockLists(db)
    seed.SeedFriendLinks(db)
}
