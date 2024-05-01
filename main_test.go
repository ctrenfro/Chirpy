package main

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestNotFound(t *testing.T) {
	s := httptest.NewServer(nil)
	t.Cleanup(s.Close)

	req, err := http.NewRequest(http.MethodGet, filepath.Join(s.URL, "index.html"), nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal((err))
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected 404, got %d", res.StatusCode)
	}
}
