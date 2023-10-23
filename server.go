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
	"github.com/gofiber/fiber/v2/middleware/cors"
)


type UserTest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	database.ConnectDb();
	database.DB.AutoMigrate(&models.User{},&models.UserSession{}, &models.Image{}, &models.UserLike{}, &models.UserMatch{}, &models.Chat{},&models.Message{})
	
	// db.AutoMigrate(&User{}, &Product{}, &Order{})

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
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

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
	routes.InteractionRoutes(app)

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