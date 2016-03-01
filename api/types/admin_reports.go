package types

import "time"

// Report holds a report header fields
type Report struct {
	ID            string       `json:"id" header:"REPORT ID"`
	Year          int          `json:"year" header:"YEAR"`
	Month         time.Month   `json:"month" header:"MONTH"`
	StartTime     time.Time    `json:"start_time" header:"START TIME"`
	EndTime       time.Time    `json:"end_time" header:"END TIME"`
	ServerSeconds float32      `json:"server_seconds" header:"SERVER TIME" show:"minifySeconds"`
	Closed        bool         `json:"closed" header:"CLOSED"`
	Lines         []Lines      `json:"lines" header:"LINES" show:"nolist"`
	AccountGroup  AccountGroup `json:"account_group" header:"ACCOUNT_GROUP" show:"nolist"`
}

// Lines holds data for report lines
type Lines struct {
	ID               string    `json:"_id" header:"ID"`
	CommissionedAt   time.Time `json:"commissioned_at" header:"COMMISSIONED_AT"`
	DecommissionedAt time.Time `json:"decommissioned_at" header:"DECOMMISSIONED_AT"`
	InstanceID       string    `json:"instance_id" header:"INSTANCE_ID"`
	InstanceName     string    `json:"instance_name" header:"INSTANCE_NAME"`
	InstanceFQDN     string    `json:"instance_fqdn" header:"INSTANCE_FQDN"`
	Consumption      float32   `json:"consumption" header:"CONSUMPTION"`
}

// AccountGroup hods account group data
type AccountGroup struct {
	ID   string `json:"_id" header:"ID"`
	Name string `json:"name" header:"NAME"`
}
