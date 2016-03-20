package types

import (
	"time"
)

type LicenseeReport struct {
	Id            string     `json:"id" header:"ID"`
	Year          int        `json:"year" header:"YEAR"`
	Month         time.Month `json:"month" header:"MONTH"`
	StartTime     time.Time  `json:"start_time" header:"START_TIME"`
	EndTime       time.Time  `json:"end_time" header:"END_TIME"`
	ServerSeconds float32    `json:"server_seconds" header:"SERVER_SECONDS"`
	Closed        bool       `json:"closed" header:"CLOSED"`
	Lines         []Lines    `json:"lines" header:"LINES"`
}
