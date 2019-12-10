package app

import (
	"github.com/mgw2007/golang-microservices/src/api/controllers/repositories"

	"github.com/mgw2007/golang-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)

}
