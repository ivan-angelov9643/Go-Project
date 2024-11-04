package main

import (
	"awesomeProject/todo-app/configuration"
	"awesomeProject/todo-app/global"
	"awesomeProject/todo-app/handlers"
	"awesomeProject/todo-app/managers/implementations"
	"awesomeProject/todo-app/middlewares"
	"awesomeProject/todo-app/middlewares/validation"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StartWebServer(config *configuration.Config) {
	router := mux.NewRouter()

	DefineRoutes(router)

	log.Info("[StartWebServer] Starting web server...")
	err := http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		log.Fatal("[StartWebServer] ListenAndServe: ", err)
	}
}

func DefineRoutes(router *mux.Router) {
	log.Info("[DefineRoutes] Defining routes...")
	router.Use(middlewares.SetJSONMiddleware)

	//Validation middleware for each one of the objects that validates input.
	//Unit tests for all routes including positive and negative tests verifying content and also middlewares

	tagHandler := handlers.NewTagHandler(implementations.NewTagManager())
	itemHandler := handlers.NewItemHandler(implementations.NewItemManager())
	listHandler := handlers.NewListHandler(implementations.NewListManager())

	router.HandleFunc("/api/tags", tagHandler.GetAll).Methods(http.MethodGet)
	router.Handle("/api/tags", validation.ValidateTagMiddleware(http.HandlerFunc(tagHandler.Create))).Methods(http.MethodPost)
	//router.HandleFunc("/api/tags", tagHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/tags/{name:"+global.TagNameRegex+"}", tagHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/tags/{name:"+global.TagNameRegex+"}", tagHandler.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/items", itemHandler.GetAll).Methods(http.MethodGet)
	router.Handle("/api/items", validation.ValidateItemMiddleware(http.HandlerFunc(itemHandler.Create))).Methods(http.MethodPost)
	//router.HandleFunc("/api/items", itemHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/items/{id:"+global.UuidRegex+"}", itemHandler.Get).Methods(http.MethodGet)
	router.Handle("/api/items/{id:"+global.UuidRegex+"}", validation.ValidateItemMiddleware(http.HandlerFunc(itemHandler.Update))).Methods(http.MethodPut)
	//router.HandleFunc("/api/items/{id:"+global.UuidRegex+"}", itemHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/items/{id:"+global.UuidRegex+"}", itemHandler.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/lists", listHandler.GetAll).Methods(http.MethodGet)
	router.Handle("/api/lists", validation.ValidateListMiddleware(http.HandlerFunc(listHandler.Create))).Methods(http.MethodPost)
	//router.HandleFunc("/api/lists", listHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/lists/{id:"+global.UuidRegex+"}", listHandler.Get).Methods(http.MethodGet)
	router.Handle("/api/lists/{id:"+global.UuidRegex+"}", validation.ValidateListMiddleware(http.HandlerFunc(listHandler.Update))).Methods(http.MethodPut)
	//router.HandleFunc("/api/lists/{id:"+global.UuidRegex+"}", listHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/lists/{id:"+global.UuidRegex+"}", listHandler.Delete).Methods(http.MethodDelete)
}

func main() {
	log.Info("[main] Starting app...")

	config, err := configuration.LoadConfig(".")
	if err != nil {
		log.Fatal("[main] Cannot load configuration: ", err)
	}

	StartWebServer(config)
}
