package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	colStartTime     = 0
	colEndTime       = 1
	colAverageWeight = 12
)

const (
	aggregationsFilepath   = "Takeout/Fit/Daily Aggregations/"
	aggregationsFileSuffix = ".csv"
	aggregationsDateFormat = "2006-01-02"
	aggregationsTimeFormat = "15:04:05.999-07:00"
)

type Record struct {
	Start    time.Time
	End      time.Time
	WeightKG float32
}

type RecordHandler func(record *Record)

func main() {
	println(os.Args[1])

	handleArchive(os.Args[1], func(record *Record) {
		fmt.Println(record)
	})
}

func handleArchive(archivePath string, handler RecordHandler) {
	archiveReader, err := zip.OpenReader(archivePath)
	if err != nil {
		panic(err)
	}
	defer archiveReader.Close()

	for _, zipFile := range archiveReader.File {
		if !strings.HasPrefix(zipFile.Name, aggregationsFilepath) {
			continue
		}
		handleDailyAggregation(zipFile, handler)
	}
}

func handleDailyAggregation(zipFile *zip.File, handler RecordHandler) {
	dateString := strings.TrimSuffix(zipFile.FileInfo().Name(), aggregationsFileSuffix)
	date, err := time.Parse(aggregationsDateFormat, dateString)
	if err != nil {
		return
	}

	fileReader, err := zipFile.Open()
	if err != nil {
		panic(err)
	}
	defer fileReader.Close()

	csvReader := csv.NewReader(fileReader)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		handleRecordLine(record, date, handler)
	}
}

func handleRecordLine(record []string, date time.Time, handler RecordHandler) {
	startTime, err := time.Parse(aggregationsTimeFormat, record[colStartTime])
	if err != nil {
		return
	}
	startTime = startTime.AddDate(date.Year(), int(date.Month())-1, date.Day()-1)

	endTime, err := time.Parse(aggregationsTimeFormat, record[colEndTime])
	if err != nil {
		return
	}
	endTime = endTime.AddDate(date.Year(), int(date.Month())-1, date.Day()-1)

	weight, err := strconv.ParseFloat(record[colAverageWeight], 32)
	if err != nil {
		return
	}

	handler(&Record{
		Start:    startTime,
		End:      endTime,
		WeightKG: float32(weight),
	})
}
