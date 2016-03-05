package types

type CloudProvider struct {
	Id                  string   `json:"id" header:"ID"`
	Name                string   `json:"name" header:"NAME"`
	RequiredCredentials []string `json:"required_credentials" header:"REQUIRED_CREDENTIALS"`
	ProvidedServices    []string `json:"provided_services" header:"PROVIDED_SERVICES"`
}
