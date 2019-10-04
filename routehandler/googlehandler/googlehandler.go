package googlehandler

import (
	"encoding/json"
	"fmt"
	"gPhotosToFlickr/config"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
)

var AppConfig *config.Config

const (
	routePrefix         = "/api/google"
	photoReadScope      = "https://www.googleapis.com/auth/drive.photos.readonly"
	redirectUriAuth     = "http://localhost:1337/api/google/authcallback"
	redirectUriToken    = "http://localhost:1337/api/google/tokencallback"
	googleTokenEndpoint = "https://oauth2.googleapis.com/token"
)

// RegisterRoutes sets up routing for the /google route prefix
func RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix(routePrefix).Subrouter()
	subRouter.HandleFunc("/health", health).Methods("GET")
	subRouter.HandleFunc("/auth", authenticate).Methods("GET")
	subRouter.HandleFunc("/authcallback", authCallback).Methods("GET")
	subRouter.HandleFunc("/tokencallback", tokenCallback).Methods("GET", "POST")

	fmt.Println("Successfully registered google routes.")
}

// health writes with a generic response message to let
// users know the API is working
func health(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.WriteHeader(200)

	msgBytes := []byte("It works!")
	_, err := respWriter.Write(msgBytes)
	if err != nil {
		fmt.Print(err.Error())
	}
}

// authenticate redirects the other to the Google OAuth login
// screen
func authenticate(respWriter http.ResponseWriter, req *http.Request) {
	oAuthUrl := buildOAuthURL(AppConfig.Google.ClientID, photoReadScope, redirectUriAuth)
	respWriter.Header().Set("Location", oAuthUrl)
	respWriter.WriteHeader(302)
}

// authCallback attempts to request a Google API token, and
// is hit when the user gives access to their Google Photos library
// via health()
func authCallback(respWriter http.ResponseWriter, req *http.Request) {
	fmt.Println("Received code!")

	code := req.FormValue("code")
	if code == "" {
		http.Error(respWriter, "code not found.", 400)
		return
	}

	// Build the URL to Google's token resolution URL based on the code
	//buildTokenURL := buildRequestTokenURL(
	//	code,
	//	AppConfig.Google.ClientID,
	//	AppConfig.Google.ClientSecret,
	//	redirectUriToken,
	//	"authorization_code")

	requestToken(code, respWriter)
}

func requestToken(code string, respWriter http.ResponseWriter) {
	tokenUrlValues := url.Values{}
	tokenUrlValues.Set("code", code)
	tokenUrlValues.Set("client_id", AppConfig.Google.ClientID)
	tokenUrlValues.Set("client_secret", AppConfig.Google.ClientSecret)
	tokenUrlValues.Set("redirect_uri", redirectUriToken)
	tokenUrlValues.Set("grant_type", "authorization_token")

	fmt.Println("Requesting auth token...")

	resp, _ := http.PostForm(googleTokenEndpoint, tokenUrlValues)

	if resp.StatusCode != 200 {
		respWriter.Header().Set("Content-Type", "application/json")

		googleErr, err := getGoogleErrorFromResponseBody(resp.Body)
		if err != nil {
			http.Error(respWriter, "something went wrong", 500)
			fmt.Println(err.Error())
			return
		}

		http.Error(respWriter, googleErr.ErrorDescription, 500)

		return
	}
}

func getGoogleErrorFromResponseBody(reader io.ReadCloser) (*GoogleError, error) {
	googleErr := new(GoogleError)
	errBytes, _ := ioutil.ReadAll(reader)

	// Attempt to decode the response body as GoogleError struct
	unmarshalErr := json.Unmarshal(errBytes, *googleErr)
	if unmarshalErr != nil {
		return googleErr, unmarshalErr
	}

	return googleErr, nil
}

// tokenCallback is called when Google sends the user back
// to this API with an auth token response. This function handles
// the response.
func tokenCallback(respWriter http.ResponseWriter, req *http.Request) {
	// TODO: Simply checking 200 isn't good enough. Should check for a range of successes
	if req.Response.StatusCode != 200 {
		// TODO: properly report on what went wrong
		http.Error(respWriter, "error occurred while getting google auth token", 500)
		return
	}

	authToken := NewAuthToken(req, time.Now())

	fmt.Println("Generated new auth token!")

	_ = json.NewEncoder(respWriter).Encode(authToken)
}

// buildOAuthURL builds and returns the Google OAuth screen URL
func buildOAuthURL(clientId string, scope string, redirectUri string) string {
	return "https://accounts.google.com/o/oauth2/v2/auth?" +
		"client_id=" + clientId +
		"&response_type=code" +
		"&scope=" + scope +
		"&redirect_uri=" + redirectUri +
		"&prompt=select_account"
}

// buildRequestTokenURL builds and returns the Google request token URL
func buildRequestTokenURL(code string, clientId string, clientSecret string, redirectUri string, grantType string) string {
	return "https://oauth2.googleapis.com/token?" +
		"code=" + code +
		"&client_id=" + clientId +
		"&client_secret=" + clientSecret +
		"&redirect_uri=" + redirectUri +
		"&grant_type=" + grantType
}
