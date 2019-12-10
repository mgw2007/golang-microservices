package main

import (
	"log"

	"github.com/mgw2007/golang-microservices/src/api/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.StartApp()
}
