package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//The call to get the IP address doesn't work.
func TestMainRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(mainRoute)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code, expected %v but got %v", http.StatusOK, rr.Code)
	}
}
