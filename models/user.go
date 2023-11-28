package models

import (
	"gorm.io/gorm"
	// "time"
)

type User struct {
	gorm.Model
	// ID        	uint           `gorm:"primaryKey" json:"id"`
    // CreatedAt 	time.Time      `json:"created_at"`
	FirstName 	 string			`json:"first_name"`
	LastName 	 string			`json:"last_name"`
	Username	 string			`gorm:"unique" json:"username"`
	Email 		 string			`gorm:"unique" json:"email"`
	Password 	 string			`json:"password"`
	UserSessions []UserSession  `json:"user_sessions"`
	// UserChats  	 []Chat 		`gorm:"foreignKey:User1ID" json:"user_chats"`
	UserLikes	 []UserLike		`json:"user_likes"`
	UserMatches	 []UserMatch	`json:"user_matches"`
	LastLocationX float64		`json:"last_location_x"`
	LastLocationY float64		`json:"last_location_y"`
}

// type User struct {
// 	gorm.Model
// 	FirstName 	 string			`json:"first_name"`
// 	LastName 	 string			`json:"last_name"`
// 	Email 		 string			`gorm:"unique" json:"email"`
// 	Password 	 string			`json:"password"`
// 	UserSessions []UserSession  `json:"user_sessions"`
// 	UserChats  	 []Chat 		`gorm:"foreignKey:User1ID" json:"user_chats"`
// 	UserLikes	 []UserLike		`json:"user_likes"`
// 	UserMatches	 []UserMatch	`json:"user_matches"`
// 	LastLocationX float64		`json:"last_location_x"`
// 	LastLocationY float64		`json:"last_location_y"`
// }