package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ObjectHandler struct {
	Text string
}

func (handler *ObjectHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if len(handler.Text) == 0 {
		log.Error("Object handler text string is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	_, err := w.Write([]byte(handler.Text))
	log.Debug("\"" + handler.Text + "\" written in response writer")
	if err != nil {
		log.Error("Error writing response: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
