package main

import (
	"awesomeProject/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/api/object1", &handlers.ObjectHandler{Text: "object1"})
	http.Handle("/api/object2", &handlers.ObjectHandler{Text: "object2"})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
