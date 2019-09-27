package googleHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const routePrefix = "/api/google"

// TODO: These need to be moved to environment variables
const googleClientID = "670357560287-i14cmjb033n19ssa2m0nng419uknqk59.apps.googleusercontent.com"
const photoReadScope = "https://www.googleapis.com/auth/drive.photos.readonly"
const redirectURI = "http://localhost:1337/api/google/callback"

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
	respWriter.Write(msgBytes)
}

// authenticate redirects the other to the Google OAuth login
// screen
func authenticate(respWriter http.ResponseWriter, req *http.Request) {
	oAuthUrl := buildOAuthURL(googleClientID, photoReadScope, redirectURI)
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

	err := json.NewEncoder(respWriter).Encode(req)
	if err != nil {
		fmt.Println(err)
		http.Error(respWriter, "could not encode response.", 500)
		return
	}
}

// buildOAuthURL builds and returns the Google OAuth screen URL
func buildOAuthURL(clientId string, scope string, redirectUri string) string {
	return "https://accounts.google.com/o/oauth2/v2/auth?client_id=" +
		googleClientID + "&response_type=code&scope=" + scope + "&redirect_uri=" + redirectUri
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
