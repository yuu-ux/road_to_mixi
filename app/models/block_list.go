package models

type BlockList struct {
    ID int64
    User1ID int
    User2ID int
	User1   User `gorm:"foreignKey:User1ID;references:UserID"`
	User2   User `gorm:"foreignKey:User2ID;references:UserID"`
}
