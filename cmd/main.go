package main

import (
	"fmt"
	// "os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	// "github.com/claudeus123/DIST2-BACKEND/ws"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/routes"

	// "github.com/claudeus123/DIST2-BACKEND/controllers"
	"github.com/claudeus123/DIST2-BACKEND/middlewares"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type UserTest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	database.ConnectDb()
	database.DB.AutoMigrate(&models.User{}, &models.UserSession{}, &models.Image{}, &models.UserLike{}, &models.UserMatch{}, &models.Chat{}, &models.Message{})

	// db.AutoMigrate(&User{}, &Product{}, &Order{})
	// chat := fiber.New(fiber.Config{
	// 	AppName: "CHAT Distribuidas II",
	// })
	app := fiber.New(fiber.Config{
		AppName: "Backend Distribuidas II",
	})
	// app.Use(func(c *fiber.Ctx) error {
	// 	// Permite solicitudes desde cualquier origen con los métodos HTTP especificados
	// 	c.Set("Access-Control-Allow-Origin", "*")
	// 	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// 	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 	// Continúa con el manejo de la solicitud
	// 	return c.Next()
	// })

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization, Content-Type",
		ExposeHeaders: "Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Use("/static", middlewares.Validate)
	app.Static("/static", "./uploads")

	routes.AuthRoutes(app)
	routes.ImageRoutes(app)
	routes.GoogleRoutes(app)
	routes.InteractionRoutes(app)
	routes.UsersRoutes(app)
	// app.Get("/session", controllers.GetSession)
	// app.Get("/users", routes.GetUsers)
	// app.Get("/users/:id", routes.GetUser)
	// app.Post("/users", routes.CreateUser)

	// app.Static("/static/", "./") util (?)

	log.Info("Hello world")
	fmt.Println("Hello world")
	app.Listen(":3333")
	// chat.Listen(":8080")
}
