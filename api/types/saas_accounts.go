package types

type SaasAccount struct {
	Id         string `json:"id" header:"ID"`
	SaasProvId string `json:"saas_provider_id" header:"SAAS PROVIDER ID"`
}

type SaasRequiredCredentials interface{}
