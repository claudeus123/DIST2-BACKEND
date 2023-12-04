package controllers

import (
	"fmt"

	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/interfaces"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/utils"
	"github.com/gofiber/fiber/v2"
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
	fmt.Println("Funcion get user")
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
		fmt.Println("usuario no encontrado?")
		return context.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	var matches []models.UserMatch
	if err := database.DB.Where("user_id = ? OR match_user_id = ?", user.ID, user.ID).Find(&matches).Error; err != nil {
		fmt.Println("Error al obtener los matches:", err)
		return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	data := interfaces.UserData{
		Id:           user.ID,

		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserSessions: user.UserSessions,
		UserLikes:    user.UserLikes,
		UserMatches:  matches,
		Gender:       user.Gender,
		Age:          user.Age,
		Bio:          user.Bio,
		Prefers:      user.Prefers,
	}
	//fmt.Println(data)
	//fmt.Println(data.UserMatches)
	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    data,
	})
}

func UserData(id uint) (interfaces.UserData, error) {
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return interfaces.UserData{}, err
	}

	if err := database.DB.Preload("UserSessions").Preload("UserLikes").Preload("UserMatches").First(&user).Error; err != nil {
		return interfaces.UserData{}, err
	}

	data := interfaces.UserData{
		Id:           user.ID,
		Email:		  user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserSessions: user.UserSessions,
		UserLikes:    user.UserLikes,
		UserMatches:  user.UserMatches,
		Gender:       user.Gender,
		Age:          user.Age,
		Bio:          user.Bio,
		Prefers:      user.Prefers,
		ImageURL:     user.ImageURL,
	}
	return data, nil
}

func GetUserDataByToken(context *fiber.Ctx) error {
	id, err := utils.GetIDFromToken(context)
	if err != nil {
		return err
	}

	data, err := UserData(uint(id))
	if err != nil {
		return err
	}

	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    data,
	})
}

func EditProfile(context *fiber.Ctx) error {
	// Obtener el token del encabezado
	token := context.Get("Authorization")
	if token == "" {
		// Manejar el caso en el que no se proporciona un token
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token not provided",
		})
	}

	// Pasar el contexto con el token a la función GetIDFromToken
	id, err := utils.GetIDFromToken(context)
	if err != nil {
		// Manejar el error, por ejemplo, devolver un error al cliente o registrar el error.
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var body struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Gender    string `json:"gender"`
		Age       int    `json:"age"`
		Bio       string `json:"bio"`
		Prefers   string `json:"prefers"`
	}
	if err := context.BodyParser(&body); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	user.Username = body.Username
	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Gender = body.Gender
	user.Age = body.Age
	user.Bio = body.Bio
	user.Prefers = body.Prefers
	database.DB.Save(&user)

	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
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
