package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestShortenValidURL(t *testing.T) {

	// Define the test request payload
	payload := []byte(`{"original": "https://www.google.com"}`)

	// Send a POST request to create a shortened URL
	res, err := http.Post("http://localhost:8082/shorten", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	// Check that the response status code is 200 OK
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	// Read the response body to get the shortened URL
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Check that the response body contains a shortened URL
	if string(body) == "" {
		t.Error("Expected a shortened URL in the response body, but got an empty string")
	}
}
