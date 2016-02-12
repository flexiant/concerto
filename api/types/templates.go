package types

import (
	"encoding/json"
)

type Template struct {
	Id                      string           `json:"id,omitempty"`
	Name                    string           `json:"name,omitempty"`
	GenericImgId            string           `json:"generic_image_id,omitempty"`
	ServiceList             []string         `json:"service_list,omitempty"`
	ConfigurationAttributes *json.RawMessage `json:"configuration_attributes,omitempty"`
}

type TemplateScript struct {
	Id               string          `json:"id"`
	Type             string          `json:"type"`
	Template_Id      string          `json:"template_id"`
	Script_Id        string          `json:"script_id"`
	Parameter_Values json.RawMessage `json:"parameter_values"`
	Execution_Order  int             `json:"execution_order"`
}

type TemplateServer struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Fqdn           string `json:"fqdn"`
	State          string `json:"state"`
	Public_ip      string `json:"public_ip"`
	Workspace_id   string `json:"workspace_id"`
	Template_id    string `json:"template_id"`
	Server_plan_id string `json:"server_plan_id"`
	Ssh_profile_id string `json:"ssh_profile_id"`
}
