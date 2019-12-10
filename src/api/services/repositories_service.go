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
}

var (
	//RepositoryService service
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
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
