package main

import (
	"fmt"
	"os"

	"github.com/g-harel/fit/internal/sources"
	"github.com/g-harel/fit/internal/sources/takeout"
)

func main() {
	println(os.Args[1])

	takeout.ReadArchive(os.Args[1], func(record *sources.Record) {
		fmt.Println(record)
	})
}
