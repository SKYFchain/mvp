package model

type Point struct {
	ObjectType string  `json:"docType"`
	Id         int64   `json:"id"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
}
