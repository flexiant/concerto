package types

type SaasProvider struct {
	Id                    string   `json:"id" header:"ID"`
	Name                  string   `json:"name" header:"NAME"`
	Required_account_data []string `json:"required_account_data" header:"REQUIRED_ACCOUNT_DATA"`
}
