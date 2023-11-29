package routes

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/claudeus123/DIST2-BACKEND/models"
	// "github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/controllers"
	// "fmt"
	// "github.com/gofiber/fiber/v2/log"
)


func UsersRoutes(app *fiber.App)  {
	middleware := app.Group("/u")
	// middleware.Use(?)
	middleware.Get("/users", controllers.GetUsers)
	middleware.Get("/users/:id", controllers.GetUser)
	middleware.Get("/profile", controllers.GetUserDataByToken)
	middleware.Patch("/editProfile", controllers.EditProfile)
	// app.Post("/users", CreateUser)
	// app.Delete("/users/:id", DeleteUser)
}

