package api

import (
	"github.com/g-harel/fit/internal/sources"
)

func getAuth() string {
	// https://developers.google.com/fit/rest/v1/authorization
	return ""
}

func createDataSource(auth string) string {
	// Check if exists
	// Create if not
	// https://developers.google.com/fit/rest/v1/data-sources
	return ""
}

func addRecordToDataset(auth string, dataSourceId string, record *sources.Record) {
	// https://developers.google.com/fit/rest/v1/datasets
	// https://stackoverflow.com/questions/36997303/submit-weight-information-distance-through-the-google-fit-rest-api
}
