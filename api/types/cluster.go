package types

// Cluster represents a Cluster entry
type Cluster struct {
	Id                string   `json:"id" header:"ID"`
	Name              string   `json:"name" header:"NAME"`
	State             string   `json:"state" header:"STATE"`
	MasterCount       int      `json:"master_count" header:"MASTER_COUNT"`
	SlaveCount        int      `json:"slave_count" header:"SLAVE_COUNT"`
	WorkspaceId       string   `json:"workspace_id" header:"WORKSPACE_ID"`
	FirewallProfileId string   `json:"firewall_profile_id" header:"FIREWALL_PROFILE_ID"`
	MasterTemplateId  string   `json:"master_template_id" header:"MASTER_TEMPLATE_ID"`
	SlaveTemplateId   string   `json:"slave_template_id" header:"SLAVE_TEMPLATE_ID"`
	Masters           []string `json:"masters" header:"MASTERS"`
}
