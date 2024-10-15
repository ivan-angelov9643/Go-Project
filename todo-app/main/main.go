package main

import (
	"awesomeProject/todo-app/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.Handle("/api/object1", &handlers.ObjectHandler{Text: "object1"})
	router.Handle("/api/object2", &handlers.ObjectHandler{Text: "object2"})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
