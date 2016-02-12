package types

type Service struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Public      bool     `json:"public"`
	License     string   `json:"license"`
	Recipes     []string `json:"recipes"`
}
