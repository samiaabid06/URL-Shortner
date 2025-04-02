package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("hasher: ", hasher)
	data := hasher.Sum(nil)
	fmt.Println("Hasher data: ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("EnodeToString: ", hash)
	fmt.Println("Short URL: ", hash[:8])
	return hash[:8]
}

func createURL(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  OriginalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	shortURL := createURL(data.URL)
	//  fmt.Fprint(w, shortURL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	// fmt.Println("Hello, World!")
	// OriginalURL := "https://www.google.com"
	// generateShortURL(OriginalURL)

	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortURLHandler)

	fmt.Println("Starting the server on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}
