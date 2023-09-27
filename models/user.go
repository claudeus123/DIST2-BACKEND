package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName 	 string			`json:"first_name"`
	LastName 	 string			`json:"last_name"`
	Email 		 string			`gorm:"unique" json:"email"`
	Password 	 string			`json:"password"`
	UserSessions []UserSession  `json:"user_sessions"`
	
}