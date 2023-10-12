package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ChatID      uint   `json:"chat_id"`
	SenderID    uint   `json:"sender_id"`
	Content     string `json:"content"`
	Sender      User   `gorm:"foreignKey:SenderID" json:"sender"`
}

// type Message struct {
// 	gorm.Model
// 	ChatID    uint   `json:"chat_id"`
// 	SenderID  uint   `json:"sender_id"`
// 	Content   string `json:"content"`
// 	// Otras propiedades del mensaje si las tienes
// }