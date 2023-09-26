package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"fmt"
	// "github.com/golang-jwt/jwt/v5"
	// "github.com/gofiber/fiber/v2/log"
)

func Register(context *fiber.Ctx) error {
	var body struct {
		// FirstName string `json:"firstName"`
		// LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	if context.BodyParser(&body) != nil {
		return context.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	// Generar un hash de la contraseña para almacenarla de manera segura
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		return context.Status(500).JSON(fiber.Map{"message": "Internal server error"})
	}
	// fmt.Println(string(hash))

	user := models.User{
		Email: body.Email,
		Password: string(hash),
	}
	fmt.Println(user)
	database.DB.Create(&user)
	fmt.Println(user)
	return context.Status(201).JSON(fiber.Map{"message": "User created", "data": user})
}