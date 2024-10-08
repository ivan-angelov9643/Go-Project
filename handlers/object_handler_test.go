package handlers_test

import (
	"awesomeProject/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObjectHandlerOK1(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/object1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := &handlers.ObjectHandler{Text: "object1"}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "object1"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestObjectHandlerOK2(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/object2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := &handlers.ObjectHandler{Text: "object2"}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "object2"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestObjectHandlerEmptyString1(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/object1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := &handlers.ObjectHandler{Text: ""}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestObjectHandlerEmptyString2(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/object2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := &handlers.ObjectHandler{Text: ""}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
