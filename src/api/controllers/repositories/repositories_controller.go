package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgw2007/golang-microservices/src/api/clients/restclient"
	"github.com/mgw2007/golang-microservices/src/api/domain/repositories"
	"github.com/mgw2007/golang-microservices/src/api/services"
	"github.com/mgw2007/golang-microservices/src/api/utils/errors"
)

//CreateRepo for CreateRepo
func CreateRepo(APIClient restclient.APIClient, c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status(), apiError)
		return
	}
	result, err := services.RepositoryService.CreateRepo(APIClient, request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//CreateManyRepos for create many in same time
func CreateManyRepos(APIClient restclient.APIClient, c *gin.Context) {
	var requests []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&requests); err != nil {
		apiError := errors.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status(), apiError)
		return
	}
	result, err := services.RepositoryService.CreateManyReposChan(APIClient, requests)
	if len(err) == 0 {
		c.JSON(http.StatusCreated, result)
	} else if len(err) == len(requests) {
		c.JSON(err[0].Status(), err)
	} else {
		// partial success
		results := struct {
			Result interface{}
			Error  interface{}
		}{result, err}

		c.JSON(http.StatusPartialContent, results)
	}
}
