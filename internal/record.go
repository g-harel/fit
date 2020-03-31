package internal

import (
	"time"
)

type RecordHandler func(record *Record)

type Record struct {
	Timestamp time.Time
	WeightKG  float32
}
