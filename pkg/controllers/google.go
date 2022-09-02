package controllers

import (
	"context"
	"fmt"
	"github.com/ssaiganesh/go-book-management-system/pkg/config"
	"io/ioutil"
	"net/http"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request) {

	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")

	// redirect to google login page
	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {

	// state
	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Fprintln(res, "states do not match")

		return
	}

	// code
	code := req.URL.Query()["code"][0]

	// configuration
	googleConfig := config.SetupConfig()

	// exchange code for token
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(res, "Code Token Exchange Failed!")
	}

	// use google api to get user info
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Fprintln(res, "User Data fetch failed")
	}

	// parse response
	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(res, "JSON Parsing Failed")
	}

	// send response to client
	fmt.Fprintln(res, string(userData))

}
