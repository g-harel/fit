package main

import (
	"fmt"
	"os"

	"github.com/g-harel/fit/internal/input/takeout"
	"github.com/g-harel/fit/internal/output/api"
	"github.com/g-harel/fit/internal/output/json"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fit <takeout zip file> <auth code>")
		os.Exit(1)
	}

	takeoutZipFile := os.Args[1]
	authCode := os.Args[2]

	jsonOutput := json.Handler{}
	takeout.ReadArchive(takeoutZipFile, jsonOutput.Handle)
	println(jsonOutput.String())

	apiOutput := api.Handler{}
	takeout.ReadArchive(takeoutZipFile, apiOutput.Handle)
	apiOutput.Write(authCode)
}
