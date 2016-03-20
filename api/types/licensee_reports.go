package types

import (
	"time"
)

type LicenseeReport struct {
	Id            string     `json:"id"`
	Year          int        `json:"year"`
	Month         time.Month `json:"month"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
	ServerSeconds float32    `json:"server_seconds"`
	Closed        bool       `json:"closed"`
	Lines         []Lines    `json:"lines"`
}
