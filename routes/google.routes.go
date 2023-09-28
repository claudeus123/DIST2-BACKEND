package routes

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/claudeus123/DIST2-BACKEND/models"
	// "github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/controllers"
	// "fmt"
	// "github.com/gofiber/fiber/v2/log"
)

func GoogleRoutes(app *fiber.App)  {
	
	// middleware.Use(?)
	app.Get("/auth/google", controllers.GoogleLogin)
	app.Get("/auth/google/callback", controllers.GoogleCallback)
	// app.Post("/users", CreateUser)
	// app.Delete("/users/:id", DeleteUser)
}