package github

//ErrorResponse for response error
type ErrorResponse struct {
	StatusCode      int     `json:"status_code"`
	Message         string  `json:"message"`
	DocumentationUr string  `json:"documentation_url"`
	Errors          []Error `json:"errors"`
}

func (r ErrorResponse) Error() string {
	return r.Message
}

//Error  for response error
type Error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
