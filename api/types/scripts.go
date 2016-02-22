package types

type Script struct {
	Id          string   `json:"id" header:"ID"`
	Name        string   `json:"name" header:"NAME"`
	Description string   `json:"description" header:"DESCRIPTION"`
	Code        string   `json:"code" header:"CODE" show:"nolist"`
	Parameters  []string `json:"parameters" header:"PARAMETERS"`
}
