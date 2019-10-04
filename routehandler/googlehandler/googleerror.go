package googlehandler

// GoogleError represents a Google API error response body
type GoogleError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
