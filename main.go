package main

const (
	regexMonetaryValue = `\$(\d{1,3},)?\d{1,3}(\.\d{2})? raised`
	regexJustStarted   = `JUST STARTED`
	mainPageURL        = "https://www.charterchallenge.ca/"
	gcpProjectID       = "public-datasets-363301"
)

// Main function for the job. For error handling, panics without returning an error because there
// is no client to return an error message to. This code is meant to be executed on a schedule
// automatically by an automated system, not manually by a user.
func main() {
	client, err := bigQueryClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	donated, goal, url, err := donatedSoFar()
	if err != nil {
		panic(err)
	}

	if err = insertRow(client, donated, goal, url); err != nil {
		panic(err)
	}
}
