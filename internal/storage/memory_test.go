package storage

import "testing"

func TestMemoryStore_SaveAndGet(t *testing.T) {
	store := NewMemoryStore()

	// Test saving
	err := store.Save("abc123", "https://google.com")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test getting
	url, found := store.Get("abc123")
	if !found {
		t.Errorf("Expected to find the URL")
	}
	if url != "https://google.com" {
		t.Errorf("Expected 'https://google.com', got '%s'", url)
	}
}

func TestMemoryStore_DuplicateError(t *testing.T) {
	store := NewMemoryStore()

	// Save first time
	store.Save("abc123", "https://google.com")

	// Save again with same key - should error
	err := store.Save("abc123", "https://different.com")
	if err == nil {
		t.Errorf("Expected an error for duplicate key, got nil")
	}
}

func TestMemoryStore_NotFound(t *testing.T) {
	store := NewMemoryStore()

	_, found := store.Get("nonexistent")
	if found {
		t.Errorf("Expected not to find a nonexistent key")
	}
}
