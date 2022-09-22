package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// Fetches data from the URL.
func fetchData() (string, error) {
	req, _ := http.NewRequest(http.MethodGet, URL, nil)
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

// Parses the amount donated so far from a web page.
func parseDonatedSoFar(webPage string) (string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", fmt.Errorf("error compiling regex expression: %w; regex string used: %s", err, regex)
	}
	matched := re.FindString(webPage)
	if len(matched) == 0 {
		return "", fmt.Errorf("after parsing with regex, matched string was of zero length")
	}
	split := strings.Split(matched, " ")
	if len(split) != 2 {
		return "", fmt.Errorf("after splitting matched string, string slice was not the expected value; string slice: %+v", split)
	}
	return split[0], nil
}

// Fetches and returns the amount donated so far as a formatted string.
func donatedSoFar() (string, error) {
	data, err := fetchData()
	if err != nil {
		return "", fmt.Errorf("could not fetch data: %w", err)
	}
	donated, err := parseDonatedSoFar(data)
	if err != nil {
		return "", fmt.Errorf("could not parse donated so far from web page data: %w", err)
	}
	return donated, nil
}
