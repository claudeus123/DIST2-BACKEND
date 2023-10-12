package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	User1ID   uint      `json:"user1_id"`
	User2ID   uint      `json:"user2_id"`
	Messages  []Message `gorm:"foreignKey:ChatID" json:"messages"`
	// User1     User      `gorm:"foreignKey:User1ID" json:"user1"`
	// User2     User      `gorm:"foreignKey:User2ID" json:"user2"`
}


// type Chat struct {
// 	gorm.Model
// 	UserId1		uint		`json:"user_id1"`
// 	UserId2		uint 		`json:"user_id2"`
	
// }