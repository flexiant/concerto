package types

type CloudAccount struct {
	Id          string `json:"id" header:"ID"`
	CloudProvId string `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
}

type RequiredCredentials interface{}
