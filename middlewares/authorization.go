package middlewares

import (
	// "os"

	"github.com/gofiber/fiber/v2"
	"github.com/claudeus123/DIST2-BACKEND/controllers"
	// "github.com/claudeus123/DIST2-BACKEND/models"
	// "github.com/golang-jwt/jwt"
	"fmt"
)

func Validate(context *fiber.Ctx) error{
	fmt.Println("Validate")
	token := context.Cookies("Authorization")
	if token == "" {
		return context.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	err := controllers.GetSession(context)
	fmt.Println(err)
	if err != nil {
		return context.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	} 
	
	return context.Next()
	// return context.Status(200).JSON(fiber.Map{
	// 	"status":  "success",
	// 	"message": "Authorized",
	// })

	// err := AuthenticateToken(token)
	// if err != nil {
	// 	return c.Status(401).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "Unauthorized",
	// 	})
	// } else {
	// 	return c.Status(200).JSON(fiber.Map{
	// 		"status":  "success",
	// 		"message": "Authorized",
	// 	})
	// }

}