package types

import (
	"encoding/json"
)

type Template struct {
	Id                      string           `json:"id,omitempty" header:"ID"`
	Name                    string           `json:"name,omitempty" header:"NAME"`
	GenericImgId            string           `json:"generic_image_id,omitempty" header:"GENERIC_IMAGE_ID"`
	ServiceList             []string         `json:"service_list,omitempty" header:"SERVICE_LIST"`
	ConfigurationAttributes *json.RawMessage `json:"configuration_attributes,omitempty" header:"CONFIGURATION_ATTRIBUTES"`
}

type TemplateScript struct {
	Id               string          `json:"id" header:"ID"`
	Type             string          `json:"type" header:"TYPE"`
	Template_Id      string          `json:"template_id" header:"TEMPLATE_ID"`
	Script_Id        string          `json:"script_id" header:"SCRIPT_ID"`
	Parameter_Values json.RawMessage `json:"parameter_values" header:"PARAMETER_VALUES"`
	Execution_Order  int             `json:"execution_order" header:"EXECUTION_ORDER"`
}

type TemplateServer struct {
	Id             string `json:"id"  header:"ID"`
	Name           string `json:"name" header:"NAME"`
	Fqdn           string `json:"fqdn" header:"FQDN"`
	State          string `json:"state" header:"STATE"`
	Public_ip      string `json:"public_ip" header:"PUBLIC_IP"`
	Workspace_id   string `json:"workspace_id" header:"WORKSPACE_ID"`
	Template_id    string `json:"template_id" header:"TEMPLATE_ID"`
	Server_plan_id string `json:"server_plan_id" header:"SERVER_PLAN_ID"`
	Ssh_profile_id string `json:"ssh_profile_id" header:"SSH_PROFILE_ID"`
}
