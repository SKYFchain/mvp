package model

import (
	"github.com/google/uuid"
	"time"
)

type Drone struct {
	ObjectType     string     `json:"docType"`
	Id             int64      `json:"id"`
	Name           string     `json:"name,omitempty"`
	Model          string     `json:"model,omitempty"`
	Capacity       int        `json:"capacity,omitempty"`
	Image          string     `json:"image,omitempty"`
	Description    string     `json:"description,omitempty"`
	Operator       string     `json:"operator,omitempty"`
	Docs           []Doc      `json:"docs,omitempty"`
	Status         string     `json:"status,omitempty"`
	NextInspection *time.Time `json:"nextInspection,omitempty"`
	UID            uuid.UUID  `json:"uid,omitempty"`
	Version        *Version   `json:"version,omitempty"`
	Point          *int64     `json:"point,omitempty"`
	ETA            *time.Time `json:"eta,omitempty"`
	Notes          string     `json:"notes,omitempty"`
}

type Version struct {
	Hardware string `json:"hardware,omitempty"`
	Software string `json:"software,omitempty"`
}

type Doc struct {
	Name string `json:"name,omitempty"`
}
