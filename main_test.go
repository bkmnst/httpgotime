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

	if !strings.Contains(body, "<title>World Clocks</title>") {
		t.Error("response does not contain HTML title")
	}

	count := strings.Count(body, "<div class=\"city\">")
	if count != 10 {
		t.Errorf("expected 10 clocks, got %d", count)
	}
}
