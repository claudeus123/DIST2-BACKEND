package middlewares

import (
	// "fmt"

	"github.com/claudeus123/DIST2-BACKEND/controllers"
	"github.com/gofiber/fiber/v2"
)

func Validate(context *fiber.Ctx) error {
	// fmt.Println("Validate")

	// Obtén el token desde el encabezado en lugar de las cookies
	authHeader := context.Get("Authorization")
	if authHeader == "" {
		// fmt.Println("Validate")
		return context.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// El valor del encabezado "Authorization" a menudo tiene un formato como "Bearer <token>"
	// Puedes dividir el valor para obtener solo el token
	// Asegúrate de manejar los casos en los que el valor no sigue este formato
	// Aquí se hace una suposición simple
	token := authHeader[len("Bearer "):]

	err := controllers.GetSessionWithToken(context, token)
	// fmt.Println(err)
	if err != nil {
		return context.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	return context.Next()
}
