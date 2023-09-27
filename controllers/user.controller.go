package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"fmt"
	// "github.com/gofiber/fiber/v2/log"
)

type Data struct {
	Id uint `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	UserSessions []models.UserSession `json:"user_sessions"`
}

func GetUsers(context *fiber.Ctx) error {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
        return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }
	
	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    users,
	})
}

func GetUser(context *fiber.Ctx) error {
	id := context.Params("id")
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	database.DB.Model(&user).Association("UserSessions").Find(&user.UserSessions)
	fmt.Println(user)
	if user.ID == 0 {
		return context.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	data := Data{
		Id: user.ID,
		Email: user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
		UserSessions: user.UserSessions,
	}
	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    data,
	})
}

// func CreateUser (context *fiber.Ctx) error {
// 	user := new(models.User)

// 	if err := context.BodyParser(user); err != nil {
// 		return context.Status(400).SendString(err.Error())
// 	}

// 	database.DB.Create(&user)
// 	fmt.Println(user)
// 	return context.SendStatus(201)
// }