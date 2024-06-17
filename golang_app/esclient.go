package main

import (
	"os"

	elastic "github.com/olivere/elastic/v7"
)

func GetESClient() (*elastic.Client, error) {
	esURL := os.Getenv("ELASTICSEARCH_URL")
	if esURL == "" {
		esURL = "http://elasticsearch:9200"
	}

	client, err := elastic.NewClient(
		elastic.SetURL(esURL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)

	return client, err
}
