package models

import (
	"gorm.io/gorm"
)


type UserLike struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	LikeUserID	 uint   `json:"target_user_id"`
	// LikeUser  	 User   `gorm:"foreignKey:LikeUserID" json:"target_user"`
}



// type UserLike struct {
// 	gorm.Model
// 	UserId				uint		`json:"user_id"`
// 	LikeUserId			uint 		`json:"like_user_id"`
// }