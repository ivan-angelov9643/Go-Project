package main

import (
	"awesomeProject/todo-app/configuration"
	"awesomeProject/todo-app/handlers"
	"awesomeProject/todo-app/middlewares"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StartWebServer(config *configuration.Config) {
	router := mux.NewRouter()
	router.Use(middlewares.SetJSONMiddleware)
	router.Handle("/api/object1", middlewares.AuthorizationMiddleware(&handlers.ObjectHandler{Text: "object1"}))
	router.Handle("/api/object2", &handlers.ObjectHandler{Text: "object2"})

	log.Info("Starting web server...")
	err := http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	log.Info("Starting app...")

	config, err := configuration.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration: ", err)
	}

	StartWebServer(config)
}
