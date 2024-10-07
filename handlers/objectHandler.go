package handlers

import "net/http"

type ObjectHandler struct {
	Text string
}

func (handler *ObjectHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte(handler.Text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
