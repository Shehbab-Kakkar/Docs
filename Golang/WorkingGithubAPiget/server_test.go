package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOctocatGists(t *testing.T) {
	server := NewServer()
	req := httptest.NewRequest(http.MethodGet, "/octocat", nil)
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var result []interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &result); err != nil {
		t.Fatalf("response is not valid JSON array: %v", err)
	}

	if len(result) == 0 {
		t.Log("warning: octocat returned zero gists (still valid)")
	}
}
