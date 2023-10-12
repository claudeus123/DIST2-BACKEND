package routes

import (
	"github.com/claudeus123/DIST2-BACKEND/controllers"
	"github.com/gofiber/fiber/v2"
)

func InteractionRoutes(app *fiber.App) {
	app.Post("/like", controllers.LikeUser)
	app.Post("/match", controllers.MakeMatch)
}