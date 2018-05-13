package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHelloHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	helloHandler(w,r)

	if status := w.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if b := w.Body.String(); b!= webContent {
		t.Fatalf("body = %s, want no", b)
	}
}
