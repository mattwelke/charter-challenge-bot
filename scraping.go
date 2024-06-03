package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Given the contents of the main web page, parses from them and returns the URL of the donation
// page.
func parseDonatePageURL(pageContents string) (string, error) {
	// Define a regex to find the donate link
	regex := regexp.MustCompile(`href="([^"]*charterchallenge\.ca[^"]*donate[^"]*)"`)

	// Search for matches
	matches := regex.FindStringSubmatch(pageContents)
	if len(matches) > 1 {
		return matches[1], nil
	} else {
		return "", fmt.Errorf("no Donate URL found")
	}
}

// Parses the amount donated so far from the donation campaign web page.
func parseDonatedSoFar(pageContents string) (string, error) {
	// First, look for the "JUST STARTED" text. If present, interpret this as
	// $0 donated so far.
	re := regexp.MustCompile(regexJustStarted)
	matched := re.FindString(pageContents)
	if matched == "JUST STARTED" {
		return "$0", nil
	}

	// If it's not $0 so far, use the monetary value regex to parse out the
	// amount donated and use that instead of $0.
	re = regexp.MustCompile(regexMonetaryValue)
	matched = re.FindString(pageContents)
	if len(matched) == 0 {
		return "", fmt.Errorf("after parsing with regex, matched string was of zero length")
	}
	split := strings.Split(matched, " ")
	if len(split) != 2 {
		return "", fmt.Errorf("after splitting matched string, string slice was not the expected value; string slice: %+v", split)
	}
	return split[0], nil
}

// Parses the goal of the current donation campaign from the donation campaign web page.
func parseGoal(pageContents string) (string, error) {
	re := regexp.MustCompile(`GOAL: \$([0-9,]+\.\d{2})`)
	match := re.FindStringSubmatch(pageContents)

	if len(match) > 1 {
		return fmt.Sprintf("$%s", match[1]), nil
	} else {
		return "", fmt.Errorf("goal not found")
	}
}

// Fetches and returns the amount donated so far as a formatted string. Also returns the goal of the
// current donation campaign as a formatted string. Also returns the URL that was determined to be
// the URL of the donation campaign page.
func donatedSoFar() (string, string, string, error) {
	mainPageContents, err := fetchPage(mainPageURL)
	if err != nil {
		return "", "", "", fmt.Errorf("could not fetch main page contents: %w", err)
	}

	donatePageURL, err := parseDonatePageURL(mainPageContents)
	if err != nil {
		return "", "", "", fmt.Errorf("could not parse donate page URL: %w", err)
	}

	donatePageContents, err := fetchPage(donatePageURL)
	if err != nil {
		return "", "", "", fmt.Errorf("could not fetch donate page contents: %w", err)
	}

	donatedSoFar, err := parseDonatedSoFar(donatePageContents)
	if err != nil {
		return "", "", "", fmt.Errorf("could not parse amount donated so far: %w", err)
	}

	goal, err := parseGoal(donatePageContents)
	if err != nil {
		return "", "", "", fmt.Errorf("could not parse goal: %w", err)
	}

	return donatedSoFar, goal, donatePageURL, nil
}
