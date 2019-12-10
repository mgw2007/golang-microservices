package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgw2007/golang-microservices/src/api/clients/restclient"
	"github.com/mgw2007/golang-microservices/src/api/config"
	"github.com/mgw2007/golang-microservices/src/api/controllers/repositories"

	"github.com/mgw2007/golang-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", func(c *gin.Context) {
		repositories.CreateRepo(restclient.APIClient{
			Client:  &http.Client{},
			BaseURL: config.GetGithubRepoURL(),
		}, c)
	})

}
