package takeout

import (
	"archive/zip"
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/g-harel/fit/internal"
)

const (
	colStartTime     = 0
	colAverageWeight = 12
)

const (
	aggregationsFilepath   = "Takeout/Fit/Daily Aggregations/"
	aggregationsFileSuffix = ".csv"
	aggregationsDateFormat = "2006-01-02"
	aggregationsTimeFormat = "15:04:05.999-07:00"
)

func ReadArchive(archivePath string, handler internal.RecordHandler) {
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

func handleDailyAggregation(zipFile *zip.File, handler internal.RecordHandler) {
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

func handleRecordLine(record []string, date time.Time, handler internal.RecordHandler) {
	timestamp, err := time.Parse(aggregationsTimeFormat, record[colStartTime])
	if err != nil {
		return
	}
	timestamp = timestamp.AddDate(date.Year(), int(date.Month())-1, date.Day()-1)

	weight, err := strconv.ParseFloat(record[colAverageWeight], 32)
	if err != nil {
		return
	}

	handler(&internal.Record{
		Timestamp: timestamp,
		WeightKG:  float32(weight),
	})
}
