package main

const (
	regex          = `\$(\d{1,3},)?\d{1,3}(\.\d{2})? raised`
	URL            = "https://www.charterchallenge.ca/donate_fall_2021"
	credsParamName = "gcpCredsBase64"
)

// Main function for the action. For error handling, panics without returning
// an error because there is no client to return an error message to. This code
// is meant to be executed on a schedule, not by a person.
func Main(params map[string]interface{}) map[string]interface{} {
	client, err := bqClient(params)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	donated, err := donatedSoFar()
	if err != nil {
		panic(err)
	}

	// Log info in BigQuery.
	err = insertRow(client, donated)
	if err != nil {
		panic(err)
	}

	// Return an empty dict to make OpenWhisk happy.
	msg := make(map[string]interface{})
	return msg
}
