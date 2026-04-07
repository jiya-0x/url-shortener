package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jiya-0x/url-shortener/internal/storage"
)

func TestShortenURL_Success(t *testing.T) {
	// Setup
	store := storage.NewMemoryStore()
	handler := &ShortenerHandler{Store: store}

	// Create a POST request with JSON body
	reqBody := `{"url":"https://google.com"}`
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()

	// Call the handler
	handler.ShortenURL(w, req)

	// Check status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	// Check response body
	var resp ShortenResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if len(resp.ShortCode) != 6 {
		t.Errorf("Expected short code of length 6, got %d", len(resp.ShortCode))
	}

	if resp.ShortURL != "http://localhost:8080/"+resp.ShortCode {
		t.Errorf("Expected short URL to match, got %s", resp.ShortURL)
	}
}

func TestShortenURL_InvalidMethod(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := &ShortenerHandler{Store: store}

	// Use GET instead of POST
	req := httptest.NewRequest(http.MethodGet, "/shorten", nil)
	w := httptest.NewRecorder()

	handler.ShortenURL(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}

func TestShortenURL_MissingURL(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := &ShortenerHandler{Store: store}

	reqBody := `{}` // Empty JSON
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.ShortenURL(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestRedirectURL_Success(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := &ShortenerHandler{Store: store}

	// Pre-save a URL
	shortCode := "abc123"
	originalURL := "https://google.com"
	store.Save(shortCode, originalURL)

	// Request the short URL
	req := httptest.NewRequest(http.MethodGet, "/"+shortCode, nil)
	w := httptest.NewRecorder()

	handler.RedirectURL(w, req)

	// Should be a redirect (302)
	if w.Code != http.StatusFound {
		t.Errorf("Expected status 302, got %d", w.Code)
	}

	// Check the Location header
	location := w.Header().Get("Location")
	if location != originalURL {
		t.Errorf("Expected Location header '%s', got '%s'", originalURL, location)
	}
}

func TestRedirectURL_NotFound(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := &ShortenerHandler{Store: store}

	// Request a nonexistent short code
	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	w := httptest.NewRecorder()

	handler.RedirectURL(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}
