package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/claudeus123/DIST2-BACKEND/utils"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/database"
	
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