// https://godoc.org/google.golang.org/api/fitness/v1

package api

import (
	"context"

	"github.com/g-harel/fit/internal"
	"golang.org/x/oauth2"
	"google.golang.org/api/fitness/v1"
	"google.golang.org/api/option"
)

type Handler struct {
	records []*internal.Record
}

func (h *Handler) Handle(record *internal.Record) {
	h.records = append(h.records, record)
}

func (h *Handler) Write(code string) {
	ctx := context.Background()
	config := &oauth2.Config{
		// TODO
	}

	token, err := config.Exchange(ctx, code)
	if err != nil {
		panic(err)
	}

	fitnessService, err := fitness.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))

	//TODO
	// createDataSource(token)
	// for {addRecordToDataset(token, ..., record)}
	println(fitnessService)
}

func createDataSource(auth string) string {
	// Check if exists
	// Create if not
	// https://developers.google.com/fit/rest/v1/data-input
	return ""
}

func addRecordToDataset(auth string, dataSourceId string, record *internal.Record) {
	// https://developers.google.com/fit/rest/v1/datasets
	// https://stackoverflow.com/questions/36997303/submit-weight-information-distance-through-the-google-fit-rest-api
}
