package types

type SSHProfile struct {
	Id          string `json:"id" header:"ID"`
	Name        string `json:"name" heade:"NAME"`
	Public_key  string `json:"public_key" header:"PUBLIC_KEY"`
	Private_key string `json:"private_key" header:"PRIVATE_KEY"`
}
