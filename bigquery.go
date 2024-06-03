package main

// Used https://cloud.google.com/bigquery/docs/samples/bigquery-table-insert-rows#bigquery_table_insert_rows-go
// as guidance

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
)

const (
	projectID = "*detect-project-id*"
	dataset   = "charter_challenge_bot"
	table     = "donation_statuses"
)

// Item represents a row item.
type Item struct {
	SoFar string
	Goal  string
	URL   string
}

func (i *Item) Save() (map[string]bigquery.Value, string, error) {
	return map[string]bigquery.Value{
		"date":   time.Now(),
		"url":    i.URL,
		"so_far": i.SoFar,
		"goal":   i.Goal,
	}, bigquery.NoDedupeID, nil
}

// Creates a BigQuery client.
func bigQueryClient() (*bigquery.Client, error) {
	var projectID string
	projectIDFromEnv := os.Getenv("GCP_PROJECT_ID")
	if projectIDFromEnv != "" {
		projectID = projectIDFromEnv
	} else {
		projectID = "*detect-project-id*"
	}
	client, err := bigquery.NewClient(context.Background(), projectID)
	if err != nil {
		return nil, fmt.Errorf("could not create BigQuery client: %w", err)
	}
	return client, nil
}

// Inserts the donation status data as a row to the BigQuery table.
func insertRow(client *bigquery.Client, donated, goal, url string) error {
	inserter := client.Dataset(dataset).Table(table).Inserter()
	items := []*Item{
		{
			SoFar: donated,
			Goal:  goal,
			URL:   url,
		},
	}
	if err := inserter.Put(context.Background(), items); err != nil {
		return fmt.Errorf("could not insert row into BigQuery: %w", err)
	}
	return nil
}
