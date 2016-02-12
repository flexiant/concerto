package types

type Script struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Code        string   `json:"code"`
	Parameters  []string `json:"parameters"`
}
