package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UserId 	  uint			`json:"user_id"`
	Url 	  string		`json:"url"`
}