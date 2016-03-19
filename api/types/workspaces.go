package types

type Workspace struct {
	Id                  string `json:"id" header:"ID"`
	Name                string `json:"name" header:"NAME"`
	Default             bool   `json:"default" header:"DEFAULT"`
	Domain_id           string `json:"domain_id" header:"DOMAIN_ID"`
	Ssh_profile_id      string `json:"ssh_profile_id" header:"SSH_PROFILE_ID"`
	Firewall_profile_id string `json:"firewall_profile_id" header:"FIREWALL_PROFILE_ID"`
}

type WorkspaceServer struct {
	Id             string `json:"id" header:"ID"`
	Name           string `json:"name" header:"NAME"`
	Fqdn           string `json:"fqdn" header:"FQDN"`
	State          string `json:"state" header:"STATE"`
	Public_ip      string `json:"public_ip" header:"PUBLIC_IP"`
	Workspace_id   string `json:"workspace_id" header:"WORKSPACE_ID"`
	Template_id    string `json:"template_id" header:"TEMPLATE_ID"`
	Server_plan_id string `json:"server_plan_id" header:"SERVER_PLAN_ID"`
	Ssh_profile_id string `json:"ssh_profile_id" header:"SSH_PROFILE_ID"`
}
