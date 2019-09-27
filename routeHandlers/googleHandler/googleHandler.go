package googleHandler

import (
	"fmt"
	"gPhotosToFlickr/config"
	"github.com/gorilla/mux"
	"net/http"
)

var AppConfig *config.Config

const routePrefix = "/api/google"
const photoReadScope = "https://www.googleapis.com/auth/drive.photos.readonly"

// RegisterRoutes sets up routing for the /google route prefix
func RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix(routePrefix).Subrouter()
	subRouter.HandleFunc("/health", health).Methods("GET")
	subRouter.HandleFunc("/auth", authenticate).Methods("GET")
	subRouter.HandleFunc("/callback", authCallback).Methods("GET")

	fmt.Println("Successfully registered google routes.")
}

// health writes with a generic response message to let
// users know the API is working
func health(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.WriteHeader(200)

	msgBytes := []byte("It works!")
	_, err := respWriter.Write(msgBytes)
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
}

// authenticate redirects the other to the Google OAuth login
// screen
func authenticate(respWriter http.ResponseWriter, req *http.Request) {
	oAuthUrl := buildOAuthURL(AppConfig.Google.ClientID, photoReadScope, AppConfig.Google.RedirectURI)
	respWriter.Header().Set("Location", oAuthUrl)
	respWriter.WriteHeader(302)
}

// authCallback attempts to request a Google API token, and
// is hit when the user gives access to their Google Photos library
// via health()
func authCallback(respWriter http.ResponseWriter, req *http.Request) {
	code := req.FormValue("code")
	if code == "" {
		http.Error(respWriter, "code not found.", 400)
		return
	}

	// Build the URL to Google's token resolution URL based on the code
	buildTokenURL := buildRequestTokenURL(code, AppConfig.Google.ClientID, AppConfig.Google.ClientSecret, AppConfig.Google.RedirectURI, "authorization_token")
	fmt.Println(buildTokenURL)
}

// buildOAuthURL builds and returns the Google OAuth screen URL
func buildOAuthURL(clientId string, scope string, redirectUri string) string {
	return "https://accounts.google.com/o/oauth2/v2/auth?client_id=" +
		clientId + "&response_type=code&scope=" + scope + "&redirect_uri=" + redirectUri
}

// buildRequestTokenURL builds and returns the Google request token URL
func buildRequestTokenURL(code string, clientId string, clientSecret string, redirectUri string, grantType string) string {
	return "https://oauth2.googleapis.com/token?" +
		"code=" + code + "&" +
		"client_id=" + clientId + "&" +
		"client_secret" + clientSecret + "&" +
		"redirect_uri" + redirectUri + "&" +
		"grant_type" + grantType
}
