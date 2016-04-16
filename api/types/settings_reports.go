package types

import "time"

// SettingsReport holds a report header fields
type SettingsReport struct {
	ID            string     `json:"id" header:"REPORT ID"`
	Year          int        `json:"year" header:"YEAR"`
	Month         time.Month `json:"month" header:"MONTH"`
	StartTime     time.Time  `json:"start_time" header:"START TIME"`
	EndTime       time.Time  `json:"end_time" header:"END TIME"`
	ServerSeconds float32    `json:"server_seconds" header:"SERVER TIME" show:"minifySeconds"`
	Closed        bool       `json:"closed" header:"CLOSED"`
	Lines         []Lines    `json:"lines" header:"LINES" show:"nolist"`
}
