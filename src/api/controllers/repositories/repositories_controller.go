package repositories

import "github.com/gin-gonic/gin"

import "github.com/mgw2007/golang-microservices/src/api/domain/repositories"

import "github.com/mgw2007/golang-microservices/src/api/utils/errors"

import "github.com/mgw2007/golang-microservices/src/api/services"

import "net/http"

//CreateRepo for CreateRepo
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status(), apiError)
		return
	}
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
