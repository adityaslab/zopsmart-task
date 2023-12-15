package models

type Train struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Platform struct {
	Number int `json:"number"`
	Train  int `json:"train"`
}
