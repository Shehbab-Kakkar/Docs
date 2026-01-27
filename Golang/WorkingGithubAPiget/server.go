package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Server holds HTTP client
type Server struct {
	client *http.Client
}

// NewServer returns a http.Handler
func NewServer() http.Handler {
	s := &Server{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	return http.HandlerFunc(s.handleUserGists)
}

// handleUserGists fetches public gists from GitHub
func (s *Server) handleUserGists(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/")
	if user == "" {
		http.Error(w, "user not specified", http.StatusBadRequest)
		return
	}

	// Pagination parameters
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	perPage := r.URL.Query().Get("per_page")
	if perPage == "" {
		perPage = "5"
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/gists?page=%s&per_page=%s", user, page, perPage)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	// Optional: GitHub token authentication
	if token := r.Header.Get("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	req.Header.Set("User-Agent", "golang-gists-api")

	resp, err := s.client.Do(req)
	if err != nil {
		http.Error(w, "failed to contact GitHub", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "GitHub API error", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
