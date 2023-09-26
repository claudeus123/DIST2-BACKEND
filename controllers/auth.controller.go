package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"fmt"
	"time"
	"os"
)

func Login(context *fiber.Ctx) error {
	
	var body struct {
		// FirstName string `json:"firstName"`
		// LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	if context.BodyParser(&body) != nil {
		return context.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}
	
	var user models.User
	database.DB.Where("email = ?", body.Email).First(&user)
	if user.ID == 0 {
		return context.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return context.Status(401).JSON(fiber.Map{"message": "Incorrect password"})
	}

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(100, 0, 0)),
			Issuer:    user.Email,
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return context.Status(401).JSON(fiber.Map{
				"success": false,
				"message": "Token Expired or invalid",
			})
		}

	context.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().AddDate(100, 0, 0),
	})
	
	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Logged in",
		// "token":   tokenString,
		"data":    user,
	})

}

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
	return context.Status(201).JSON(fiber.Map{"message": "User created"})
}