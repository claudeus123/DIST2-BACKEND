package controllers

import (
	"fmt"
	"strconv"

	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/interfaces"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"golang.org/x/crypto/bcrypt"
	// "github.com/claudeus123/DIST2-BACKEND/mail"
	"github.com/jordan-wright/email"
	"net/smtp"
	// "github.com/gofiber/fiber/v2/log"
)

func GetUsers(context *fiber.Ctx) error {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    users,
	})
}

func GetUser(context *fiber.Ctx) error {
	fmt.Println("Funcion get user")
	id, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	data, err := UserData(uint(id))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	fmt.Println(data)
	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    data,
	})
	
}

func UserData(id uint) (interfaces.UserData, error) {
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return interfaces.UserData{}, err
	}

	if user.ID == 0 {
		fmt.Println("usuario no encontrado?")
		return interfaces.UserData{}, nil
	}

	if err := database.DB.Preload("UserSessions").Preload("UserLikes").First(&user).Error; err != nil {
		return interfaces.UserData{}, err
	}

	chats := []models.Chat{}
	if err := database.DB.Where("user1_id = ? OR user2_id = ?", user.ID, user.ID).Find(&chats).Error; err != nil {
		return interfaces.UserData{}, err
	}

	matches := []models.UserMatch{}
	if err := database.DB.Where("user_id = ? OR match_user_id = ?", user.ID, user.ID).Find(&matches).Error; err != nil {
		return interfaces.UserData{}, err
	}

	data := interfaces.UserData{
		Id:           user.ID,
		Email:		  user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserSessions: user.UserSessions,
		UserLikes:    user.UserLikes,
		UserMatches:  matches,
		Gender:       user.Gender,
		Age:          user.Age,
		Bio:          user.Bio,
		Prefers:      user.Prefers,
		ImageURL:     user.ImageURL,
		UserChats:    chats,
		Username:	  user.Username,
		Latitude: 	  user.Latitude,
		Longitude:    user.Longitude,
	}
	return data, nil
}

func GetUserDataByToken(context *fiber.Ctx) error {
	id, err := utils.GetIDFromToken(context)
	if err != nil {
		return err
	}

	data, err := UserData(uint(id))
	if err != nil {
		return err
	}

	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    data,
	})
}

func EditProfile(context *fiber.Ctx) error {
	// Obtener el token del encabezado
	token := context.Get("Authorization")
	if token == "" {
		// Manejar el caso en el que no se proporciona un token
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token not provided",
		})
	}

	// Pasar el contexto con el token a la función GetIDFromToken
	id, err := utils.GetIDFromToken(context)
	if err != nil {
		// Manejar el error, por ejemplo, devolver un error al cliente o registrar el error.
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var body struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Gender    string `json:"gender"`
		Age       int    `json:"age"`
		Bio       string `json:"bio"`
		Prefers   string `json:"prefers"`
	}
	if err := context.BodyParser(&body); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	user.Username = body.Username
	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Gender = body.Gender
	user.Age = body.Age
	user.Bio = body.Bio
	user.Prefers = body.Prefers
	database.DB.Save(&user)

	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
	})
}

func Forgot (context *fiber.Ctx) error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando variables de entorno: %v", err)
	}
	
	smtpAuthAddress := os.Getenv("SMTP_AUTH_ADDRESS")
	smtpServerAddress := os.Getenv("SMTP_SERVER_ADDRESS")
	emailSenderName := os.Getenv("EMAIL_SENDER_NAME")
	emailSenderAddress := os.Getenv("EMAIL_SENDER_ADDRESS")
	emailSenderPassword := os.Getenv("EMAIL_SENDER_PASSWORD")
	var body struct {
		Email string `json:"email"`
	}
	if err := context.BodyParser(&body); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		
		return context.Status(400).JSON(fiber.Map{"message": "User not found"})
	}
	password, err := utils.GeneratePassword(20)
	if err != nil {
		return context.Status(500).JSON(fiber.Map{"message": "Internal server error"})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return context.Status(500).JSON(fiber.Map{"message": "Internal server error"})
	}
	e := email.NewEmail()
		e.From = emailSenderName + " <" + emailSenderAddress + ">"
		e.To = []string{body.Email}
		e.Subject = "Recuperación de contraseña"
		// e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<p>La nueva contraseña es: <strong>" + password +"</strong></p>")
		e.Send(smtpServerAddress, smtp.PlainAuth("", emailSenderAddress, emailSenderPassword, smtpAuthAddress))

		// return context.Status(200).JSON(fiber.Map{"message": "Email sent"})
	
	user.Password = string(hash)
	database.DB.Save(&user)

	return context.Status(200).JSON(fiber.Map{"message": "Email sent"})
}


func ChangePassword(context *fiber.Ctx) error {
	id, err := utils.GetIDFromToken(context)
	if err != nil {
		return err
	}
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	var body struct {
		Password string `json:"password"`
	}
	if err := context.BodyParser(&body); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return context.Status(500).JSON(fiber.Map{"message": "Internal server error"})
	}
	user.Password = string(hash)
	database.DB.Save(&user)

	Logout(context)
	return context.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
	})
}

func SetLocation(context *fiber.Ctx) error {
	userID, err := utils.GetIDFromToken(context)
	if err != nil {
		return context.Status(500).JSON(fiber.Map{"message": "Internal server error"})
	}

	var body struct {
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	if err := context.BodyParser(&body); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return context.Status(500).JSON(fiber.Map{"message": "Internal server error"})
	}

	user.Latitude = body.Latitude
	user.Longitude = body.Longitude
	database.DB.Save(&user)

	return context.Status(200).JSON(fiber.Map{"message": "Success", "data": user})
}
// func CreateUser (context *fiber.Ctx) error {
// 	user := new(models.User)

// 	if err := context.BodyParser(user); err != nil {
// 		return context.Status(400).SendString(err.Error())
// 	}

// 	database.DB.Create(&user)
// 	fmt.Println(user)
// 	return context.SendStatus(201)
// }
