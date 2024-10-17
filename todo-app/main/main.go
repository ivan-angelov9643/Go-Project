package main

import (
	"awesomeProject/todo-app/configuration"
	"awesomeProject/todo-app/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.Info("Starting app...")
	config, err := configuration.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration: ", err)
	}

	router := mux.NewRouter()

	router.Handle("/api/object1", &handlers.ObjectHandler{Text: "object1"})
	router.Handle("/api/object2", &handlers.ObjectHandler{Text: "object2"})

	log.Info("Starting web server")
	err = http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
