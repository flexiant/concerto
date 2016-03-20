package types

type FirewallProfile struct {
	Id          string `json:"id" header:"ID"`
	Name        string `json:"name,omitempty" header:"NAME"`
	Description string `json:"description,omitempty" header:"DESCRIPTION"`
	Default     bool   `json:"default,omitempty" header:"DEFAULT"`
	Rules       []Rule `json:"rules,omitempty" header:"RULES"`
}

type Rule struct {
	Protocol string `json:"ip_protocol" header:"IP_PROTOCOL"`
	MinPort  int    `json:"min_port" header:"MIN_PORT"`
	MaxPort  int    `json:"max_port" header:"MAX_PORT"`
	CidrIp   string `json:"source" header:"SOURCE"`
}
