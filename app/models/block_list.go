package models

type BlockList struct {
	ID      int64 `json:"id"`
	User1ID int   `json:"user1_id" gorm:"index"`
	User2ID int   `json:"user2_id" gorm:"index"`
	User1   User  `gorm:"foreignKey:User1ID;references:UserID"`
	User2   User  `gorm:"foreignKey:User2ID;references:UserID"`
}
