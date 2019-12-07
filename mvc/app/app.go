package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp func its init point for runining the app
func StartApp() {
	mapUrls()
	log.Fatal(router.Run(":5050"))
}
