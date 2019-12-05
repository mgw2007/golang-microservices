package app

import "net/http"

import "github.com/mgw2007/golang-microservices/mvc/controllers"

import "log"

// StartApp func its init point for runining the app
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	log.Fatal(http.ListenAndServe(":5050", nil))
}
