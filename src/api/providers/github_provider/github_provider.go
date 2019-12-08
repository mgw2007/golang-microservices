package github_provider

import "github.com/mgw2007/golang-microservices/src/api/domain/github"

import "fmt"

import "github.com/mgw2007/golang-microservices/src/api/clients/restclient"

import "net/http"

import "log"

import "encoding/json"

import "io/ioutil"

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "Token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(token string) string {
	return fmt.Sprintf(headerAuthorizationFormat, token)
}

//CreateRepo for CreateRepo
func CreateRepo(APIClient restclient.APIClient, token string, req github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(token))
	response, err := APIClient.Post(req, headers)
	if err != nil {
		log.Printf("error when trying create repo in github: %v\n", err)
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Printf("error when trying read response body in github: %v\n", err)
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}
	if response.StatusCode > 299 {
		var errorResponse github.ErrorResponse
		if err := json.Unmarshal(bytes, &errorResponse); err != nil {
			log.Printf("error when Unmarshal error response body from github: %v\n", err)
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response error body",
			}
		}
		errorResponse.StatusCode = response.StatusCode
		return nil, &errorResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when Unmarshal success response body from github: %v\n", err)
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid json response success body",
		}
	}
	return &result, nil
}
