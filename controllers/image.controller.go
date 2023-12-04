package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	// "strings"
	"fmt"
	// "strconv"
	// "os"
	// "github.com/google/uuid"
	"github.com/claudeus123/DIST2-BACKEND/database"
	"github.com/claudeus123/DIST2-BACKEND/models"
	"github.com/claudeus123/DIST2-BACKEND/utils"
)

func ImageUpload(context *fiber.Ctx) error {
	userID, err := utils.GetIDFromToken(context)
    file, err := context.FormFile("image")

    if err != nil {
        log.Println("image upload error --> ", err)
        return context.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

    }

	// orderStr := context.FormValue("order")
	// order, err := strconv.Atoi(orderStr)
	// if err != nil {
	// 	return context.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	// }

	// isProfileStr := context.FormValue("is_profile")
	// isProfile, err := strconv.ParseBool(isProfileStr)
	// if err != nil {
	// 	return context.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	// }

    // // generate new uuid for image name 
    // uniqueId := uuid.New()

    // // remove "- from imageName"

    // filename := strings.Replace(uniqueId.String(), "-", "", -1)

    // extract image extension from original file filename

    

    // save image to ./images dir 
    err = context.SaveFile(file, fmt.Sprintf("./uploads/%d.jpg", userID))
	//Verificar si existe la img
    if err != nil {
        log.Println("image save error --> ", err)
        return context.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
    }

    // generate image url to serve to client using CDN

    // imageUrl := fmt.Sprintf("http://localhost:3000/images/%s", image)

    // create meta data and send to client

    data := map[string]interface{}{
        "imageName": fmt.Sprintf("./uploads/%d.jpg", userID),
        "header":    file.Header,
        "size":      file.Size,
    }

	// var imageDB models.Image
	// imageDB.Url = fmt.Sprintf("./uploads/%d-%s.jpg", userID, orderStr )
	// imageDB.UserId = uint(userID)
	// imageDB.Order = order
	// imageDB.IsProfile = isProfile

	var user models.User
	result := database.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return context.JSON(fiber.Map{"status": 404, "message": "User not found", "data": nil})
	}
	user.ImageURL = fmt.Sprintf("/static/%d.jpg", userID)
	database.DB.Save(&user)

    return context.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func ImageServe(context *fiber.Ctx) error{
	userID, err := utils.GetIDFromToken(context)
	if err != nil {
		return context.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

// 	var imageDB models.Image
// 	result := database.DB.Where("user_id = ? AND id = ?", userID, imageID).First(&imageDB)
// 	if result.Error != nil {
// 		return context.JSON(fiber.Map{"status": 404, "message": "Image not found", "data": nil})
// 	}

// 	return context.SendFile(imageDB.Url)

// 	// Obtener ID de los parámetros de la ruta
	// user, err := UserData(uint(userID))
	// if err != nil {
	// 	return context.JSON(fiber.Map{"status": 404, "message": "User not found", "data": nil})
	// }

	// imagePath := "./uploads/" + fmt.Sprint(user.Id) + "-" + imageID + ".jpg"
	// if _, err := os.Stat(imagePath); os.IsNotExist(err) {
	// 				// Si la imagen no existe, devolver un error 404
	// 				return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 					"error": "Imagen no encontrada.",
	// 				})
	// }

	var user models.User
	result := database.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return context.JSON(fiber.Map{"status": 404, "message": "User not found", "data": nil})
	}
	
		// Estructura de respuesta JSON
		response := fiber.Map{
			"user":  user,
			"image": user.ImageURL,
		}

		return context.JSON(response)
	// return nil
}