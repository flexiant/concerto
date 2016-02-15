package types

import (
	"encoding/json"
)

// Template stores blueprint templates
type Template struct {
	ID                      string          `json:"id,omitempty" header:"ID"`
	Name                    string          `json:"name,omitempty" header:"NAME"`
	GenericImgID            string          `json:"generic_image_id,omitempty" header:"GENERIC IMAGE ID"`
	ServiceList             []string        `json:"service_list,omitempty" header:"SERVICE LIST" show:"nolist"`
	ConfigurationAttributes json.RawMessage `json:"configuration_attributes,omitempty" header:"CONFIGURATION ATTRIBUTES" show:"nolist"`
}

// TemplateScript stores a templates' script info
type TemplateScript struct {
	ID              string          `json:"id" header:"ID"`
	Type            string          `json:"type" header:"TYPE"`
	ExecutionOrder  int             `json:"execution_order" header:"EXECUTION ORDER"`
	TemplateID      string          `json:"template_id" header:"TEMPLATE ID"`
	ScriptID        string          `json:"script_id" header:"SCRIPT ID"`
	ParameterValues json.RawMessage `json:"parameter_values" header:"PARAMETER VALUES"`
}

// TemplateServer stores servers associated with the template
type TemplateServer struct {
	ID           string `json:"id"  header:"ID"`
	Name         string `json:"name" header:"NAME"`
	Fqdn         string `json:"fqdn" header:"FQDN"`
	State        string `json:"state" header:"STATE"`
	PublicIP     string `json:"public_ip" header:"PUBLIC IP"`
	WorkspaceID  string `json:"workspace_id" header:"WORKSPACE ID"`
	TemplateID   string `json:"template_id" header:"TEMPLATE ID"`
	ServerPlanID string `json:"server_plan_id" header:"SERVER PLAN ID"`
	SSHProfileID string `json:"ssh_profile_id" header:"SSH PROFILE ID"`
}

// TemplateScriptCredentials stores credentials to servers
type TemplateScriptCredentials interface{}
