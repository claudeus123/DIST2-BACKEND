package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/claudeus123/DIST2-BACKEND/utils"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/database"

	"github.com/joho/godotenv"
	"log"
	"fmt"
	"os"
	"bytes"
    "encoding/json"
    // "fmt"
    // "log"
    "net/http"
)

func LikeUser(context *fiber.Ctx) error {
	var body struct {
		UserId uint `json:"user_id"`
	}

	if context.BodyParser(&body) != nil {
		return context.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	userId, err := utils.GetIDFromToken(context)
	if err != nil {
		return context.JSON(fiber.Map{"status": 401, "message": "Unauthorized"})
	}

	var like models.UserLike
	like.UserID = uint(userId)
	like.LikeUserID = uint(body.UserId)

	database.DB.Create(&like)


	return context.JSON(fiber.Map{"status": 200, "message": "success", "data": like})

}

func CheckMatch(context *fiber.Ctx) bool {
	var body struct {
		UserId uint `json:"user_id"`
	}

	if context.BodyParser(&body) != nil {
		return false
	}

	userId, err := utils.GetIDFromToken(context)
	if err != nil {
		return false
	}

	userData, err := UserData(body.UserId)
	if err != nil {
		return false
	}

	for _, like := range userData.UserLikes {
		if like.LikeUserID == uint(userId) {
			return true
		}
	}
	return false

}

func MakeMatch(context *fiber.Ctx) error {
	var check bool = CheckMatch(context)

	var body struct {
		UserId uint `json:"user_id"`
	}

	// var err error
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error cargando variables de entorno: %v", err)
    }

	wsUrl := os.Getenv("WS_URL")
	
	if context.BodyParser(&body) != nil {
		return context.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	if check {
		id, err := utils.GetIDFromToken(context)
		if err != nil {
			return context.JSON(fiber.Map{"status": 401, "message": "Unauthorized"})
		}

		var match models.UserMatch
		match.UserID = uint(id)
		match.MatchUserID = uint(body.UserId)

		database.DB.Create(&match)

		// CREAR CHAT
		var chat models.Chat
		chat.User1ID = uint(id)
		chat.User2ID = uint(body.UserId)
		database.DB.Create(&chat)

		//Crear CHAT EN WEBSOCKET
		var chatBody struct {
			ChatId uint `json:"chat_id"`
			User1Id uint `json:"user1_id"`
			User2Id uint `json:"user2_id"`
		}
		chatBody.ChatId = chat.ID
		chatBody.User1Id = uint(id)
		chatBody.User2Id = uint(body.UserId)

		values := chatBody
		jsonData, err := json.Marshal(values)

		if err != nil {
			log.Fatal(err)
		}

		postRoute := wsUrl + "/ws/createChat"
		resp, err := http.Post(postRoute, "application/json",
			bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
		fmt.Println(res["json"])

		return context.JSON(fiber.Map{"match": "true"})
	} else {
		return context.Status(200).JSON(fiber.Map{"match": "false"})
	}

	
	// var match models.UserMatch
	// match.UserID = uint(userId)
	// match.MatchUserID = uint(body.UserId)

	// database.DB.Create(&match)

	// return context.JSON(fiber.Map{"status": 200, "message": "success", "data": match})

}