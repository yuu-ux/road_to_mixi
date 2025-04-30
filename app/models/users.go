package models

type User struct {
	ID     int64  `json:"id"`
	UserID int    `json:"user_id" gorm:"unique"`
	Name   string `json:"name"`
}
