package auth

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig oauth2.Config
}

type GoogleUserInfo struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
	Picture   string `json:"picture"`
}

var AppConfig Config

func GoogleOauth() oauth2.Config {
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

func GetGoogleLoginURL(state string) string {
	return AppConfig.GoogleLoginConfig.AuthCodeURL(state)
}

func GetGoogleToken(code string) (*oauth2.Token, error) {
	token, err := AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetGoogleUserInfo(accessToken string) (*GoogleUserInfo, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	userInfoURL := os.Getenv("USERINFO")

	client := AppConfig.GoogleLoginConfig.Client(context.Background(), &oauth2.Token{AccessToken: accessToken})
	resp, err := client.Get(userInfoURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user info: " + resp.Status)
	}

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
