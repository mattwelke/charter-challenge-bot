package main

import (
	"fmt"
	"io"
	"net/http"
)

// Fetches data from a URL.
func fetchPage(url string) (string, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("FROM", "https://github.com/mattwelke/charter-challenge-for-fair-voting-bot")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error during HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not get body of response: %w", err)
	}

	return string(body), nil
}
