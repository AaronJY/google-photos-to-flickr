package tests

import "testing"

func TestRegisterRoutes() {

}

func TestHealth(t *testing.T) {

}

func TestAuthenticate() {
}

func TestAuthCallback() {

}

func TestBuildOAUthUrl(t *testing.T) {
	expected := "https://accounts.google.com/o/oauth2/v2/auth?client_id=670357560287-i14cmjb033n19ssa2m0nng419uknqk59.apps.googleusercontent.com&response_type=code&scope=https://www.googleapis.com/auth/drive.photos.readonly&redirect_uri=http://localhost:1337/api/google/callback"
	if (google)
}