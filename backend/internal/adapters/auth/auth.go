package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig oauth2.Config
}

var AppConfig Config

func InitializeGoogleOAuth() oauth2.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	userInfoEmail := os.Getenv("USERINFO_EMAIL")
	userInfoProfile := os.Getenv("USERINFO_PROFILE")

	scopeList := []string{userInfoEmail, userInfoProfile}

	AppConfig.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       scopeList,
		Endpoint:     google.Endpoint,
	}
	return AppConfig.GoogleLoginConfig
}
