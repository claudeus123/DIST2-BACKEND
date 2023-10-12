package utils

import (
	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/gofiber/fiber/v2"
	"errors"
)

func GetIDFromToken(context *fiber.Ctx) (int, error) {
	var token models.UserSession
	jwt := context.Cookies("Authorization")
	database.DB.Where("token = ?", jwt).First(&token) 

	if token.Token == "" {
		return -1, errors.New("Invalid token")
	}
	return int(token.UserId), nil
}