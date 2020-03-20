package sources

import (
	"time"
)

type RecordHandler func(record *Record)

type Record struct {
	// TODO single timestamp.
	Start    time.Time
	End      time.Time
	WeightKG float32
}
