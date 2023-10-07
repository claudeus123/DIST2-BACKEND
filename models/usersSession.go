package models

import (
	"gorm.io/gorm"
)

type UserSession struct {
	gorm.Model
	UserId		uint		`json:"user_id"`
	IsValid 	bool		`json:"is_valid"`
	Token 		string		`gorm:"index" json:"token"`
}