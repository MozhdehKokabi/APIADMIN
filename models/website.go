package models

type ReqWebsite struct {
	Domain  string `json:"domain"`
	Name    string `json: "name"`
	Address string `json:"address"`
}
