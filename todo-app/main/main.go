package main

import (
	"awesomeProject/todo-app/configuration"
	"awesomeProject/todo-app/global"
	"awesomeProject/todo-app/handlers"
	"awesomeProject/todo-app/middlewares"
	"awesomeProject/todo-app/middlewares/validation"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StartWebServer(config *configuration.Config) {
	router := mux.NewRouter()

	DefineRoutes(router)

	log.Info("Starting web server...")
	err := http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func DefineRoutes(router *mux.Router) {
	router.Use(middlewares.SetJSONMiddleware)

	// Create handlers for operations of each type that are served on:
	//
	//GET on “api/%s” - Get All
	//POST on “api/%s” - Create New
	//GET on “/api/%s/ {id}” - Get existing
	//- PUT on “/api/%s/{id}
	//” - Update existing
	//
	//DELETE on “/api/%s/ {id}
	//” - Delete existing
	//
	//Handler should work with JSON as input and JSON as response using middlewares
	//All routing paths should be centrally described for easy reading and maintenance
	//Validation middleware for each one of the objects that validates input.
	//Unit tests for all routes including positive and negative tests verifying content and also middlewares

	tagHandler := handlers.NewTagHandler()
	router.HandleFunc("/api/tags", tagHandler.GetAll).Methods(http.MethodGet)
	router.Handle("/api/tags", validation.ValidateTag(http.HandlerFunc(tagHandler.Create))).Methods(http.MethodPost)
	router.HandleFunc("/api/tags/{id:"+global.UuidRegex+"}", tagHandler.Get).Methods(http.MethodGet)
	router.Handle("/api/tags/{id:"+global.UuidRegex+"}", validation.ValidateTag(http.HandlerFunc(tagHandler.Update))).Methods(http.MethodPut)
	router.HandleFunc("/api/tags/{id:"+global.UuidRegex+"}", tagHandler.Delete).Methods(http.MethodDelete)
}

func main() {
	log.Info("Starting app...")

	config, err := configuration.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration: ", err)
	}

	StartWebServer(config)
}
