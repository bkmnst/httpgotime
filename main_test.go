package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	greet(w, r)
	response := w.Result()
	defer response.Body.Close()

	body := w.Body.String()

	if !strings.Contains(body, "Hello World!") {
		t.Error("response does not contain Hello World!")
	}
}
