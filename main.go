package main

const (
	regexMonetaryValue = `\$(\d{1,3},)?\d{1,3}(\.\d{2})? raised`
	regexJustStarted   = `JUST STARTED`
	URL                = "https://www.charterchallenge.ca/donate_charter_challenge"
	credsParamName     = "gcpCredsBase64"
)

// Main function for the action. For error handling, panics without returning
// an error because there is no client to return an error message to. This code
// is meant to be executed on a schedule, not by a person.
func Main(params map[string]interface{}) map[string]interface{} {
	client, err := bigQueryClient(params)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	donated, err := donatedSoFar()
	if err != nil {
		panic(err)
	}

	if err = insertRow(client, donated); err != nil {
		panic(err)
	}

	// Return an empty dict to make OpenWhisk happy.
	msg := make(map[string]interface{})
	return msg
}
