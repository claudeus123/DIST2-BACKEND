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


func AuthRoutes(app *fiber.App)  {
	auth := app.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	// auth.Post("/google/login", controllers.GoogleAuth)
	// auth.Post("/google/register", controllers.GoogleSignup)
	auth.Get("/logout", middlewares.Validate, controllers.Logout)

	// app.Post("/users", CreateUser)
	// app.Delete("/users/:id", DeleteUser)
}

