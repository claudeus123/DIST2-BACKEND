package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"fmt"
	// "github.com/gofiber/fiber/v2/log"
)

func GetSession(context *fiber.Ctx) error {
	token := context.Cookies("Authorization")
	if token == "" {
		return context.Status(404).JSON(fiber.Map{"message": "Token not found"})
	}
	
	var session models.UserSession
	if err := database.DB.Where("token = ?", token).First(&session).Error; err != nil {
		// return context.Status(404).JSON(fiber.Map{"message": "Token not found"})
		return fmt.Errorf("Token not found")
	}
	// fmt.Println(session.Token)
	if session.IsValid == false {
		// return context.Status(401).JSON(fiber.Map{"message": "Invalid token"})
		return fmt.Errorf("Invalid token")
	}
	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    session,
	})
}