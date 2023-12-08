package interfaces

import (
	"github.com/claudeus123/DIST2-BACKEND/models"
)

type UserData struct {
	Id           uint                 `json:"id"`
	Email        string               `json:"email"`
	FirstName    string               `json:"firstName"`
	LastName     string               `json:"lastName"`
	UserSessions []models.UserSession `json:"user_sessions"`
	UserLikes    []models.UserLike    `json:"user_likes"`
	UserMatches  []models.UserMatch   `json:"user_matches"`
	UserChats    []models.Chat        `json:"user_chats"`
	Gender       string               `json:"gender"`
	Username	 string				  `json:"username"`
	Age          int                  `json:"age"`
	Bio          string               `json:"bio"`
	Prefers      string               `json:"prefers"`
	ImageURL     string               `json:"image_url"`
	Latitude float64				`json:"latitude"`
	Longitude float64				`json:"longitude"`
}

type InteractionData struct {
	Id           uint                 `json:"id"`
	Email        string               `json:"email"`
	FirstName    string               `json:"firstName"`
	LastName     string               `json:"lastName"`
	Gender       string               `json:"gender"`
	Username	 string				  `json:"username"`
	Age          int                  `json:"age"`
	Bio          string               `json:"bio"`
	Prefers      string               `json:"prefers"`
	ImageURL     string               `json:"image_url"`
}
