package google

type Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (googleErr Error) GetError() string {
	return googleErr.ErrorDescription + " - " + googleErr.Error
}
