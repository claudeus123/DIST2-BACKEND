package controllers

import (
	"fmt"

	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/gofiber/fiber/v2"
)

func GetSessionWithToken(context *fiber.Ctx, token string) error {
	if token == "" {
		return context.Status(404).JSON(fiber.Map{"message": "Token not found"})
	}

	var session models.UserSession
	if err := database.DB.Where("token = ?", token).First(&session).Error; err != nil {
		return fmt.Errorf("Token not found")
	}

	if session.IsValid == false {
		return fmt.Errorf("Invalid token")
	}

	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    session,
	})
}
