package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"fmt"
	"os"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/routes"
	// "github.com/claudeus123/DIST2-BACKEND/controllers"
	"github.com/claudeus123/DIST2-BACKEND/middlewares"
)


type UserTest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	database.ConnectDb();
	database.DB.AutoMigrate(&models.User{},&models.UserSession{}, &models.Image{})
		// database.DB.Exec("ALTER TABLE user_sessions ALTER COLUMN user_id SET DATA TYPE integer")
	// db.AutoMigrate(&User{}, &Product{}, &Order{})

    app := fiber.New(fiber.Config{
		AppName: "Backend Distribuidas II",	
	})
    app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
    })
	app.Get("/testing", func(c *fiber.Ctx) error {
		// Obtener ID de los parámetros de la ruta
		

		// Obtener datos del usuario (puedes obtener estos datos de tu base de datos)
		user := UserTest{
			ID:    1,
			Name:  "Usuario Ejemplo",
			Email: "usuario@example.com",
		}

		// Obtener la ruta de la imagen del usuario (puedes obtener esta ruta de tu base de datos u otro almacenamiento)
		imagePath := "./uploads/" + fmt.Sprint(user.ID) + ".jpg"

		// Verificar si la imagen existe
		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			// Si la imagen no existe, devolver un error 404
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Imagen no encontrada.",
			})
		}

		// Estructura de respuesta JSON
		response := fiber.Map{
			"user":  user,
			"image": "/static/" + fmt.Sprint(user.ID) + ".jpg",
		}

		return c.JSON(response)
	})
	app.Static("/static", "./uploads")
	
	routes.ImageRoutes(app)
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