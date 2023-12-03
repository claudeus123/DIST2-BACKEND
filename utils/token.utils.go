package utils

import (
	"errors"
	"fmt"

	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/gofiber/fiber/v2"
)

func GetIDFromToken(context *fiber.Ctx) (int, error) {
	var token models.UserSession

	// Obtén el token desde el encabezado en lugar de las cookies
	authHeader := context.Get("Authorization")
	if authHeader == "" {
		return -1, errors.New("Authorization header missing")
	}

	// El valor del encabezado "Authorization" a menudo tiene un formato como "Bearer <token>"
	// Puedes dividir el valor para obtener solo el token
	// Asegúrate de manejar los casos en los que el valor no sigue este formato
	// Aquí se hace una suposición simple
	tokenValue := authHeader[len("Bearer "):]

	// Imprime el token extraído para depuración
	fmt.Println("Token Value:", tokenValue)

	database.DB.Where("token = ?", tokenValue).First(&token)

	// Imprime el token desde la base de datos para depuración
	fmt.Println("Token from DB:", token.Token)

	if token.Token == "" {
		return -1, errors.New("Invalid token")
	}

	return int(token.UserId), nil
}
