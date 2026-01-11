package entity

type Monster struct {
	ID     uint     `json:"id"`
	Name   string   `json:"name"`
	Powers []string `json:"powers"`
}
