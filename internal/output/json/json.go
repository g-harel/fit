package json

import (
	"encoding/json"
	"sort"

	"github.com/g-harel/fit/internal"
)

type jsonRecord struct {
	Timestamp int64   `json:"timestamp"`
	Weight    float32 `json:"weight"`
}

type Handler struct {
	records []*internal.Record
}

func (h *Handler) Handle(record *internal.Record) {
	h.records = append(h.records, record)
}

func (h *Handler) String() string {
	sort.Sort(h)
	records := []*jsonRecord{}
	for i := 0; i < len(h.records); i++ {
		records = append(records, &jsonRecord{
			Timestamp: h.records[i].Timestamp.Unix(),
			Weight:    h.records[i].WeightKG,
		})
	}
	bytes, _ := json.Marshal(records)
	return string(bytes)
}

// Sort interface.
func (h *Handler) Len() int           { return len(h.records) }
func (h *Handler) Swap(i, j int)      { h.records[i], h.records[j] = h.records[j], h.records[i] }
func (h *Handler) Less(i, j int) bool { return h.records[i].Timestamp.Before(h.records[j].Timestamp) }
