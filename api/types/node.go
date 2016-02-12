package types

type Node struct {
	Id        string `json:"id" header:"ID"`
	Name      string `json:"name" header:"NAME"`
	Fqdn      string `json:"fqdn" header:"FQDN"`
	PublicIp  string `json:"public_ip" header:"PUBLIC_IP"`
	State     string `json:"state" header:"STATE"`
	Os        string `json:"os" header:"OS"`
	Plan      string `json:"plan" header:"PLAN"`
	FleetName string `json:"fleet_name" header:"FLEET_NAME"`
	Master    bool   `json:"is_master" header:"IS_MASTER"`
}
