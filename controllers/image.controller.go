package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
	"fmt"
	"strconv"
	"github.com/google/uuid"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/models"
)

func ImageUpload(context *fiber.Ctx) error {

    // parse incomming image file

    file, err := context.FormFile("image")

    if err != nil {
        log.Println("image upload error --> ", err)
        return context.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

    }

    // generate new uuid for image name 
    uniqueId := uuid.New()

    // remove "- from imageName"

    filename := strings.Replace(uniqueId.String(), "-", "", -1)

    // extract image extension from original file filename

    fileExt := strings.Split(file.Filename, ".")[1]

    // generate image from filename and extension
    image := fmt.Sprintf("%s.%s", filename, fileExt)

    // save image to ./images dir 
    err = context.SaveFile(file, fmt.Sprintf("./uploads/%s", image))

    if err != nil {
        log.Println("image save error --> ", err)
        return context.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
    }

    // generate image url to serve to client using CDN

    // imageUrl := fmt.Sprintf("http://localhost:3000/images/%s", image)

    // create meta data and send to client

    data := map[string]interface{}{
        "imageName": image,
        "header":    file.Header,
        "size":      file.Size,
    }

	var token models.UserSession
	jwt := context.Cookies("Authorization")
	database.DB.Where("token = ?", jwt).First(&token) 

	if token.Token == "" {
		return context.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	orderStr := context.FormValue("order")
	order, err := strconv.Atoi(orderStr)
	if err != nil {
		return context.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	isProfileStr := context.FormValue("is_profile")
	isProfile, err := strconv.ParseBool(isProfileStr)
	if err != nil {
		return context.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	var imageDB models.Image
	imageDB.Url = fmt.Sprintf("uploads/%s",image)
	imageDB.UserId = token.UserId
	imageDB.Order = order
	imageDB.IsProfile = isProfile

	database.DB.Create(&imageDB)


    return context.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}