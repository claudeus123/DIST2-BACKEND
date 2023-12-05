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


func UsersRoutes(app *fiber.App)  {
	user := app.Group("/users")

	user.Get("/", middlewares.Validate, controllers.GetUsers)
	user.Get("/get/:id", middlewares.Validate, controllers.GetUser)
	user.Get("/profile", middlewares.Validate, controllers.GetUserDataByToken)
	user.Patch("/editProfile", middlewares.Validate, controllers.EditProfile)
	user.Post("/resetPassword", controllers.Forgot)
}

