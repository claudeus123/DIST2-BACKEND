package routes

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/claudeus123/DIST2-BACKEND/models"
	// "github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/controllers"
	"github.com/claudeus123/DIST2-BACKEND/middlewares"

	// "fmt"
	// "github.com/gofiber/fiber/v2/log"
)


func ImageRoutes(app *fiber.App)  {
	image := app.Group("/image")

	image.Post("/upload", middlewares.Validate, controllers.ImageUpload)
	image.Get("/profile", middlewares.Validate, controllers.ImageServe)
	// app.Post("/users", CreateUser)
	// app.Delete("/users/:id", DeleteUser)
}

