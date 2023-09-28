package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/joho/godotenv"
	"os"
	"log"
)

func SetupGoogleConfig() *oauth2.Config{
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error cargando variables de entorno: %v", err)
    }

	conf := &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	return conf
}