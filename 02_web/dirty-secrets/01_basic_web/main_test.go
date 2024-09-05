package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllHandler(t *testing.T) {
	// 1. Test Request erzeugen
	req, _ := http.NewRequest("GET", "/api", nil)

	// 2. HTTP-Recorder erzeugen (erfüllt http.ResponseWriter)
	recorder := httptest.NewRecorder()

	// 3. Handler aufrufen
	getAllHandler(recorder, req)

	// 4. HTTP-Status code prüfen
	if recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code: got %v expected %v", recorder.Code, http.StatusOK)
	}
}
