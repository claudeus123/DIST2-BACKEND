package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
	"fmt"
	"github.com/google/uuid"
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

    return context.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}