package googlehandler

// GoogleError represents a Google API error response body
type GoogleError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// GetError returns a string representation of the full Google error
func (googleErr GoogleError) GetError() string {
	return googleErr.ErrorDescription + " - " + googleErr.Error
}
