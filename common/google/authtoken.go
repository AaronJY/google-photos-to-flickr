package google

import (
	"time"
)

type AuthToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresOn    time.Time `json:"expires_on"`
	TokenType    string    `json:"token_type"`
}
