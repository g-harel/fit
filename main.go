package main

import (
	"os"

	"github.com/g-harel/fit/internal/input/takeout"
	"github.com/g-harel/fit/internal/output/json"
)

func main() {
	println(os.Args[1])
	jsonOutput := json.Handler{}
	takeout.ReadArchive(os.Args[1], jsonOutput.Handle)
	println(jsonOutput.String())
}
