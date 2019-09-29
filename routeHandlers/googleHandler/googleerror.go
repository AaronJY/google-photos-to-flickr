package googleHandler

type GoogleError struct{
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
}