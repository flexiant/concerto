package types

type Server struct {
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

type Dns struct {
	Id        string `json:"id" header:"ID"`
	Name      string `json:"name" header:"NAME"`
	Content   string `json:"content" header:"CONTENT"`
	Type      string `json:"type" header:"TYPE"`
	IsFQDN    bool   `json:"is_fqdn" header:"IS_FQDN"`
	Domain_id string `json:"domain_id" header:"DOMAIN_ID"`
}

type ScriptChar struct {
	Id               string   `json:"id" header:"ID"`
	Type             string   `json:"type" header:"TYPE"`
	Parameter_values struct{} `json:"parameter_values" header:"PARAMETER_VALUES"`
	Template_id      string   `json:"template_id" header:"TEMPLATE_ID"`
	Script_id        string   `json:"script_id" header:"SCRIPT_ID"`
}
