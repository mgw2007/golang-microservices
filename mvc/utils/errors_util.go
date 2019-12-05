package utils

//ApplicationError for handle error
type ApplicationError struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	StatusCode int    `json:"status"`
}
