package models

import (
	"gorm.io/gorm"
)
type UserMatch struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	MatchUserID  uint   `json:"match_user_id"`
	// MatchUser    User   `gorm:"foreignKey:MatchUserID" json:"match_user"`
}
// type UserMatch struct {
// 	gorm.Model
// 	UserId				uint		`json:"user_id"`
// 	MatchUserId			uint 		`json:"match_user_id"`
// }