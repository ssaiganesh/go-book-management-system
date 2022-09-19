package config

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile("/Users/shankar.ganesh/go/src/github.com/ssaiganesh/go-book-management-system/.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     viperEnvVariable("client_id"),
		ClientSecret: viperEnvVariable("client_secret"),
		RedirectURL:  "https:localhost:9010/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/aith/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf

}
