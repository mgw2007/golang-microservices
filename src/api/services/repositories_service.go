package services

import (
	"strings"

	"github.com/mgw2007/golang-microservices/src/api/clients/restclient"
	"github.com/mgw2007/golang-microservices/src/api/config"
	"github.com/mgw2007/golang-microservices/src/api/domain/github"
	"github.com/mgw2007/golang-microservices/src/api/domain/repositories"
	"github.com/mgw2007/golang-microservices/src/api/providers/github_provider"
	"github.com/mgw2007/golang-microservices/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(APIClient restclient.APIClient, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateManyRepos(APIClient restclient.APIClient, requests []repositories.CreateRepoRequest) ([]*repositories.CreateRepoResponse, []errors.ApiError)
	CreateManyReposChan(APIClient restclient.APIClient, requests []repositories.CreateRepoRequest) ([]*repositories.CreateRepoResponse, []errors.ApiError)
}

var (
	//RepositoryService service
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}
func (s *repoService) CreateManyReposChan(APIClient restclient.APIClient, requests []repositories.CreateRepoRequest) ([]*repositories.CreateRepoResponse, []errors.ApiError) {
	type RepoRespons struct {
		Resonse *repositories.CreateRepoResponse
		Error   errors.ApiError
	}
	var responses []*repositories.CreateRepoResponse
	var errors []errors.ApiError
	output := make(chan RepoRespons)
	defer close(output)

	for _, request := range requests {
		go func(output chan RepoRespons, request repositories.CreateRepoRequest) {
			respone, err := s.CreateRepo(APIClient, request)
			if err != nil {
				output <- RepoRespons{
					Error: err,
				}
			} else {
				output <- RepoRespons{
					Resonse: respone,
				}
			}
		}(output, request)
	}
	for ok := range output {
		if ok.Resonse != nil {
			responses = append(responses, ok.Resonse)
		} else if ok.Error != nil {
			errors = append(errors, ok.Error)
		}
		if len(responses)+len(errors) == len(requests) {
			break
		}
	}
	return responses, errors
}

func (s *repoService) CreateManyRepos(APIClient restclient.APIClient, requests []repositories.CreateRepoRequest) ([]*repositories.CreateRepoResponse, []errors.ApiError) {
	var responses []*repositories.CreateRepoResponse
	var errors []errors.ApiError

	for _, request := range requests {
		respone, err := s.CreateRepo(APIClient, request)
		if err != nil {
			errors = append(errors, err)
		} else {
			responses = append(responses, respone)
		}
	}
	return responses, errors
}

func (s *repoService) CreateRepo(APIClient restclient.APIClient, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid Repository Name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	response, err := github_provider.CreateRepo(APIClient, config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewAPIError(err.StatusCode, err.Message)
	}
	return &repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}, nil
}
