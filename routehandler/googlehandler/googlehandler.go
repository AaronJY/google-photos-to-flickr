package googlehandler

import (
	"encoding/json"
	"fmt"
	"github.com/AaronJY/google-photos-to-flickr/common/google"
	"github.com/AaronJY/google-photos-to-flickr/config"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var AppConfig *config.Config

const (
	routePrefix         = "/api/google"
	photoReadScope      = "https://www.googleapis.com/auth/drive.photos.readonly"
	redirectUriAuth     = "http://localhost:1337/api/google/authcallback"
	googleTokenEndpoint = "https://oauth2.googleapis.com/token"
)

// RegisterRoutes sets up routing for the /google route prefix
func RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix(routePrefix).Subrouter()
	subRouter.HandleFunc("/health", health).Methods("GET")
	subRouter.HandleFunc("/auth", authenticate).Methods("GET")
	subRouter.HandleFunc("/authcallback", authCallback).Methods("GET")

	fmt.Println("Successfully registered google routes.")
}

// health writes with a generic response message to let users know the API is working
func health(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.WriteHeader(200)

	msgBytes := []byte("It works!")
	_, err := respWriter.Write(msgBytes)
	if err != nil {
		fmt.Print(err.Error())
	}
}

// authenticate redirects the other to the Google OAuth login screen
func authenticate(respWriter http.ResponseWriter, req *http.Request) {
	oAuthUrl := buildOAuthURL(AppConfig.Google.ClientID, photoReadScope, redirectUriAuth)
	respWriter.Header().Set("Location", oAuthUrl)
	respWriter.WriteHeader(302)
}

// authCallback attempts to request a Google API token
func authCallback(respWriter http.ResponseWriter, req *http.Request) {
	fmt.Println("Successfully retrieved Google OAuth code")

	code := req.FormValue("code")
	if code == "" {
		http.Error(respWriter, "code not found.", 400)
		return
	}

	requestToken(code, respWriter)
}

// requestToken requests a Google authentication based on an OAauth code
func requestToken(code string, respWriter http.ResponseWriter) {
	tokenUrlValues := url.Values{}
	tokenUrlValues.Set("code", code)
	tokenUrlValues.Set("client_id", AppConfig.Google.ClientID)
	tokenUrlValues.Set("client_secret", AppConfig.Google.ClientSecret)
	tokenUrlValues.Set("redirect_uri", redirectUriAuth)
	tokenUrlValues.Set("grant_type", "authorization_code")

	resp, _ := http.PostForm(googleTokenEndpoint, tokenUrlValues)

	if resp.StatusCode != 200 {
		respWriter.Header().Set("Content-Type", "application/json")

		googleErr, err := getGoogleErrorFromResponseBody(resp.Body)
		if err != nil {
			http.Error(respWriter, "something went wrong", 500)
			fmt.Println("Error retrieving Google auth token:", err.Error())
			return
		}

		http.Error(respWriter, googleErr.GetError(), 500)
		return
	}

	fmt.Println("Successfully retrieved Google auth token")

	authToken := new(google.AuthToken)
	err := json.NewDecoder(resp.Body).Decode(authToken)
	if err != nil {
		http.Error(respWriter, "could not read auth token", 500)
		fmt.Println("Could not decode auth token json:", err.Error())
	}

	respWriter.Header().Set("Location", "http://localhost:1337?googleAuth=1&googleapikey=" + authToken.AccessToken)
	respWriter.WriteHeader(302)
}

// getGoogleErrorFromResponseBody gets a Google error response string from a response body
func getGoogleErrorFromResponseBody(reader io.ReadCloser) (*google.Error, error) {
	googleErr := new(google.Error)
	errBytes, _ := ioutil.ReadAll(reader)

	defer reader.Close()

	// Attempt to decode the response body as GoogleError struct
	unmarshalErr := json.Unmarshal(errBytes, &googleErr)
	if unmarshalErr != nil {
		return googleErr, unmarshalErr
	}

	return googleErr, nil
}

// buildOAuthURL builds and returns the Google OAuth screen URL
func buildOAuthURL(clientId string, scope string, redirectUri string) string {
	return "https://accounts.google.com/o/oauth2/v2/auth?" +
		"client_id=" + clientId +
		"&response_type=code" +
		"&scope=" + scope +
		"&redirect_uri=" + redirectUri +
		"&prompt=select_account" +
		"&access_type=offline"
}
