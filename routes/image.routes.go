package routes

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/claudeus123/DIST2-BACKEND/models"
	// "github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/controllers"
	// "fmt"
	// "github.com/gofiber/fiber/v2/log"
)


func ImageRoutes(app *fiber.App)  {
	app.Post("/upload", controllers.ImageUpload)
	// app.Post("/users", CreateUser)
	// app.Delete("/users/:id", DeleteUser)
}

