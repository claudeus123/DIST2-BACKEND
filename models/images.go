package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UserId 	  uint			`json:"user_id"`
	Url 	  string		`json:"url"`
	IsProfile bool			`json:"is_profile"`
	Order 	  int			`json:"order"`
	IsRemoved bool			`json:"is_removed"`
}