// https://godoc.org/google.golang.org/api/fitness/v1
// https://github.com/a-h/gofit/
// https://github.com/1self/google-fit-integration/blob/master/src/main.go

package api

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/g-harel/fit/internal"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/fitness/v1"
	"google.golang.org/api/option"
)

type Handler struct {
	records []*internal.Record
}

func (h *Handler) Handle(record *internal.Record) {
	h.records = append(h.records, record)
}

func (h *Handler) Write(credentialsPath string) error {

	//TODO
	// createFitnessService(credentialsPath)
	// createDataSource(token)
	// for {addRecordToDataset(token, ..., record)}

	return nil
}

func createFitnessService(credentialsPath string) (*fitness.Service, error) {
	ctx := context.Background()

	config, err := loadConfigFromFile(credentialsPath, fitness.FitnessBodyWriteScope)
	if err != nil {
		return nil, fmt.Errorf("load credentials: %v", err)
	}

	// TODO
	code := "code"

	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("convert auth to token: %v", err)
	}

	fitnessService, err := fitness.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return nil, fmt.Errorf("create service: %v", err)
	}

	return fitnessService, nil
}

func loadConfigFromFile(filename string, scopes ...string) (*oauth2.Config, error) {
	json, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read secret file: %v", err)
	}

	conf, err := google.ConfigFromJSON(json)
	if err != nil {
		return nil, fmt.Errorf("config from json: %v", err)
	}

	conf.Scopes = append(conf.Scopes, scopes...)
	return conf, nil
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
