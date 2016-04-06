package types

type LoadBalancer struct {
	Id                          string `json:"id" header:"ID"`
	Name                        string `json:"name" header:"NAME"`
	Fqdn                        string `json:"fqdn" header:"FQDN"`
	Protocol                    string `json:"protocol" header:"PROTOCOL"`
	Port                        int    `json:"port" header:"PORT"`
	Algorithm                   string `json:"algorithm" header:"ALGORITHM"`
	SslCertificate              string `json:"ssl_certificate" header:"SSL_CERTIFICATE"`
	Ssl_certificate_private_key string `json:"ssl_certificate_private_key" header:"SSL_CERTIFICATE_PRIVATE_KEY"`
	Domain_id                   string `json:"domain_id" header:"DOMAIN_ID"`
	Cloud_provider_id           string `json:"cloud_provider_id" header:"CLOUD_PROVIDER_ID"`
	Traffic_in                  int    `json:"traffic_in" header:"TRAFFIC_IN"`
	Traffic_out                 int    `json:"traffic_out" header:"TRAFFIC_OUT"`
}

type LBNode struct {
	Id       string `json:"id" header:"ID"`
	Name     string `json:"name" header:"NAME"`
	PublicIp string `json:"public_ip" header:"PUBLIC_IP"`
	State    string `json:"state" header:"STATE"`
	ServerId string `json:"server_id" header:"SERVER_ID"`
	Port     int    `json:"port" header:"PORT"`
}
