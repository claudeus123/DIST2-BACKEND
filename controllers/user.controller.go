package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"fmt"
	"github.com/claudeus123/DIST2-BACKEND/interfaces"
	// "github.com/gofiber/fiber/v2/log"
)



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
	// VER TEMA DE SELECT * FROM USERS MATCH_ID1 = ID OR MATCH_ID2 = ID ALGO ASI
	if err := database.DB.Preload("UserSessions").Preload("UserLikes").Preload("UserMatches").First(&user).Error; err != nil {
		// Manejar el error, por ejemplo, devolver un error al cliente o registrar el error.
		fmt.Println("Error al cargar el usuario:", err)
		return err
	}

	if user.ID == 0 {
		return context.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	var matches []models.UserMatch
    if err := database.DB.Where("user_id = ? OR match_user_id = ?", user.ID, user.ID).Find(&matches).Error; err != nil {
        fmt.Println("Error al obtener los matches:", err)
        return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }

	data := interfaces.UserData{
		Id: user.ID,
		Email: user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
		UserSessions: user.UserSessions,
		UserLikes: user.UserLikes,
		UserMatches: matches,

	}
	fmt.Println(data)
	fmt.Println(data.UserMatches)
	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    data,
	})
}

func UserData (id uint) (interfaces.UserData, error) {
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return interfaces.UserData{}, err
	}

	if err := database.DB.Preload("UserSessions").Preload("UserLikes").Preload("UserMatches").First(&user).Error; err != nil {
		return interfaces.UserData{}, err
	}

	data := interfaces.UserData{
		Id:           user.ID,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserSessions: user.UserSessions,
		UserLikes:    user.UserLikes,
		UserMatches:  user.UserMatches,
	}
	return data, nil
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