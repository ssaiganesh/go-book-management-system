package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "test", // to update
		ClientSecret: "test", // to update
		RedirectURL:  "https:localhost:9010/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/aith/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf

}
