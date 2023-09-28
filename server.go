package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"fmt"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/routes"
	// "github.com/claudeus123/DIST2-BACKEND/controllers"
	"github.com/claudeus123/DIST2-BACKEND/middlewares"
)

func main() {

	database.ConnectDb();
	database.DB.AutoMigrate(&models.User{},&models.UserSession{})
		// database.DB.Exec("ALTER TABLE user_sessions ALTER COLUMN user_id SET DATA TYPE integer")
	// db.AutoMigrate(&User{}, &Product{}, &Order{})

    app := fiber.New(fiber.Config{
		AppName: "Backend Distribuidas II",	
	})
    app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
    })
	routes.GoogleRoutes(app)
	routes.AuthRoutes(app)
	
	app.Use(middlewares.Validate)
	routes.UsersRoutes(app)
	// app.Get("/session", controllers.GetSession)
	// app.Get("/users", routes.GetUsers)
	// app.Get("/users/:id", routes.GetUser)
	// app.Post("/users", routes.CreateUser)

	// app.Static("/static/", "./") util (?)

	log.Info("Hello world")
	fmt.Println("Hello world")
    app.Listen(":3000")
}