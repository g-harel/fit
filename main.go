package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/g-harel/fit/internal/input/takeout"
	"github.com/g-harel/fit/internal/output/api"
	"github.com/g-harel/fit/internal/output/json"
)

var takeoutArchivePath = flag.String("takeout", "", "takeout archive filename")
var credentialsPath = flag.String("credentials", "", "credentials json filename")

func main() {
	flag.Parse()

	if *takeoutArchivePath == "" || *credentialsPath == "" {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	jsonOutput := json.Handler{}
	takeout.ReadArchive(*takeoutArchivePath, jsonOutput.Handle)
	println(jsonOutput.String())

	apiOutput := api.Handler{}
	takeout.ReadArchive(*takeoutArchivePath, apiOutput.Handle)
	apiOutput.Write(*credentialsPath)
}
