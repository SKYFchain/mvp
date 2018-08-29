package model

import (
	"time"
)

type Mission struct {
	ObjectType  string    `json:"docType"`
	Id          int64     `json:"id"`
	Name        string    `json:"name,omitempty"`
	Route       int64     `json:"route,omitempty"`
	Point       int64     `json:"point,omitempty"`
	Drone       int64     `json:"drone,omitempty"`
	Cargo       string    `json:"cargo,omitempty"`
	Customer    string    `json:"customer,omitempty"`
	Certs       []Cert    `json:"certs,omitempty"`
	Legal       []Legal   `json:"legal,omitempty"`
	Price       int64     `json:"price,omitempty"`
	Status      string    `json:"status,omitempty"`
	ETA         time.Time `json:"eta,omitempty"`
	ETD         time.Time `json:"etd,omitempty"`
	Description string    `json:"description,omitempty"`
	Txns        []Txn     `json:"txns,omitempty"`
}

type Txn struct {
	Hash string `json:"hash"`
}
type Cert struct {
	Name string `json:"name,omitempty"`
}

type Legal struct {
	Name string `json:"name,omitempty"`
}
