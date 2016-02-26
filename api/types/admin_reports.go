package types

import "time"

// Report holds a report header fields
type Report struct {
	Id             string       `json:"id" header:"REPORT ID"`
	Year           int          `json:"year" header:"YEAR"`
	Month          time.Month   `json:"month" header:"MONTH"`
	Start_time     time.Time    `json:"start_time" header:"START_TIME"`
	End_time       time.Time    `json:"end_time" header:"END_TIME"`
	Server_seconds float32      `json:"server_seconds" header:"SERVER_SECONDS"`
	Closed         bool         `json:"closed" header:"CLOSED"`
	Li             []Lines      `json:"lines" header:"LINES" show:"nolist"`
	Account_group  AccountGroup `json:"account_group" header:"ACCOUNT_GROUP" show:"nolist"`
}

type Lines struct {
	Id                string    `json:"_id" header:"ID"`
	Commissioned_at   time.Time `json:"commissioned_at" header:"COMMISSIONED_AT"`
	Decommissioned_at time.Time `json:"decommissioned_at" header:"DECOMMISSIONED_AT"`
	Instance_id       string    `json:"instance_id" header:"INSTANCE_ID"`
	Instance_name     string    `json:"instance_name" header:"INSTANCE_NAME"`
	Instance_fqdn     string    `json:"instance_fqdn" header:"INSTANCE_FQDN"`
	Consumption       float32   `json:"consumption" header:"CONSUMPTION"`
}

type AccountGroup struct {
	Id   string `json:"_id" header:"ID"`
	Name string `json:"name" header:"NAME"`
}
