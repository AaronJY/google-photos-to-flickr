package googleHandler

import (
	"net/http"
	"strconv"
	"time"
)

type AuthToken struct {
	accessToken string
	refreshToken string
	expiresIn int
	expiresOn time.Time
	tokenType string
}

func NewAuthToken(response *http.Request, createdOn time.Time) *AuthToken {
	authToken := new(AuthToken)
	authToken.accessToken = response.FormValue("access_token")
	authToken.refreshToken = response.FormValue("refresh_token")

	authTokenInt, _ := strconv.Atoi(response.FormValue("expires_in"))
	authToken.expiresIn = authTokenInt

	authToken.expiresOn = createdOn.Add(time.Duration(authToken.expiresIn))

	authToken.tokenType = response.FormValue("token_type")

	return authToken
}

