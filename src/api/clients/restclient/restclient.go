package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//APIClient for set used APIClient
type APIClient struct {
	Client  *http.Client
	BaseURL string
}

//Post for send post request
func (api *APIClient) Post(body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, api.BaseURL, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := api.Client
	return client.Do(request)
}
