package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jiya-0x/url-shortener/internal/handlers"
	"github.com/jiya-0x/url-shortener/internal/storage"
)

func main() {

	store := storage.NewMemoryStore()
	fmt.Println("Storage initialized (in-memory)")

	handler := &handlers.ShortenerHandler{
		Store: store,
	}
	fmt.Println("Handlers initialized")

	http.HandleFunc("/shorten", handler.ShortenURL)

	http.HandleFunc("/", handler.RedirectURL)
	fmt.Println("Routes registered")

	fmt.Println("\n- Server starting on http://localhost:8080")
	fmt.Println("- POST to http://localhost:8080/shorten with JSON:")
	fmt.Println("  {\"url\":\"https://example.com\"}")
	fmt.Println("- Visit http://localhost:8080/{shortCode} to be redirected")
	fmt.Println("\nPress Ctrl+C to stop the server\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
