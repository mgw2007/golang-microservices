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
