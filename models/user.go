package models

import (
	"time"
)

type User struct {
	ID 			uint		`gorm:"autoIncrement" json:"id"`
	FirstName 	string		`json:"first_name"`
	LastName 	string		`json:"last_name"`
	Email 		string		`gorm:"primaryKey" json:"email"`
	Password 	string		`json:"password"`
	CreatedAt 	time.Time	`json:"created_at"`
}