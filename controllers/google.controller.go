package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/claudeus123/DIST2-BACKEND/config"
	"golang.org/x/oauth2"
	// "io/ioutil"
	"encoding/json"
	// "encoding/base64"
	// "golang.org/x/oauth2/google"
	// "github.com/claudeus123/DIST2-BACKEND/config"
	// "github.com/claudeus123/DIST2-BACKEND/models"
	// "github.com/claudeus123/DIST2-BACKEND/database"
	"fmt"
	"net/http"
	// "github.com/gofiber/fiber/v2/log"
)

func GoogleLogin(context *fiber.Ctx)  error{
	googleConfig := config.SetupGoogleConfig()
	// url := googleConfig.AuthCodeURL("not-implemetned")
	url := googleConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)

	return context.Redirect(url)
	// return context.Status(200).JSON(fiber.Map{
	// 	"success": true,
	// 	"message": "Success",
	// 	"data":    "Google Login",
	// })
}

type googleAuthResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func GoogleCallback(context *fiber.Ctx) error {
	code := context.Query("code")
	conf := config.SetupGoogleConfig()
	
	token, err := conf.Exchange(context.Context(), code)
	if err != nil {
	  return context.SendStatus(fiber.StatusInternalServerError)
	}
	fmt.Println("Este es el token mi onichan")
	// fmt.Println(token)
	fmt.Println(token.AccessToken)


	response, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return err;
	}

	defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return err;
	// }

	googleResponse := googleAuthResponse{}
	err = json.NewDecoder(response.Body).Decode(&googleResponse)
	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusInternalServerError)
		return  context.JSON(err)
	}

	
	return context.JSON(googleResponse)
	// resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", token.AccessToken))
    // if httpErr != nil {
    //     return httpErr
    // }
	// return context.SendString("Hello from google callback")
	// code := context.Query("code")
	// fmt.Println(code)
    // return context.Status(200).JSON(fiber.Map{
	// 	"success": true,
	// 	"message": "Success",
	// 	"data":    "Google Login",
	// })
}