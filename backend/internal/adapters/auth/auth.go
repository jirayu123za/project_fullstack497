package auth

import (
	"backend_fullstack/internal/config"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig oauth2.Config
}

var AppConfig Config

func InitializeGoogleOAuth() oauth2.Config {
	config.LoadEnv()
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
