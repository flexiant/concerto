package types

type ServerPlan struct {
	Id              string  `json:"id" header:"ID"`
	Name            string  `json:"name" header:"NAME"`
	Memory          int     `json:"memory" header:"MEMORY"`
	CPUs            float32 `json:"cpus" header:"CPUS"`
	Storage         int     `json:"storage" header:"STORAGE"`
	LocationId      string  `json:"location_id" header:"LOCATION_ID"`
	CloudProviderId string  `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
}
