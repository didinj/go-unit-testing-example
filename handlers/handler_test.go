package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Djamware", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	res := w.Result()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}

	expected := "Hello, Djamware!"
	if strings.TrimSpace(string(body)) != expected {
		t.Errorf("got %q, want %q", string(body), expected)
	}
}

func TestGreetHandler(t *testing.T) {
	payload := `{"name": "Gopher"}`
	req := httptest.NewRequest(http.MethodPost, "/greet", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GreetHandler(w, req)

	res := w.Result()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}

	expected := "Hello, Gopher!"
	if strings.TrimSpace(string(body)) != expected {
		t.Errorf("got %q, want %q", string(body), expected)
	}
}
