package model

type Route struct {
	ObjectType string  `json:"docType"`
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Active     bool    `json:"active"`
	Distance   int     `json:"distance"`
	Altitude   []int   `json:"altitude"`
	Points     []int64 `json:"points"`
}
