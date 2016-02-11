package types

// Domain represents a domain entry
type Domain struct {
	ID      string `json:"id" header:"ID"`
	Name    string `json:"name" header:"NAME"`
	TTL     int    `json:"ttl" header:"TTL"`
	Contact string `json:"contact" header:"CONTACT"`
	Minimum int    `json:"minimum" header:"MINIMUM"`
	Enabled bool   `json:"enabled" header:"ENABLED"`
}

// DomainRecord represents a domain record entry
type DomainRecord struct {
	ID       string `json:"id" header:"ID"`
	Type     string `json:"type" header:"TYPE"`
	Name     string `json:"name" header:"NAME"`
	Content  string `json:"content" header:"CONTENT"`
	TTL      int    `json:"ttl" header:"TTL"`
	Prio     int    `json:"prio" header:"PRIO"`
	ServerID string `json:"server_id" header:"SERVER ID"`
	DomainID string `json:"domain_id" header:"DOMAIN ID"`
}
