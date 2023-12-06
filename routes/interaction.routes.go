package routes

import (
	"github.com/claudeus123/DIST2-BACKEND/controllers"
	"github.com/claudeus123/DIST2-BACKEND/middlewares"
	"github.com/gofiber/fiber/v2"
)

func InteractionRoutes(app *fiber.App) {
	interaction := app.Group("/interaction")

	interaction.Get("/users", middlewares.Validate, controllers.GetPossibleInteractions)
	interaction.Post("/like", middlewares.Validate, controllers.LikeUser)
	interaction.Post("/match", middlewares.Validate, controllers.MakeMatch)
}