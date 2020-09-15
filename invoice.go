package main

type Invoice struct {
	Customer     string `json:"customer"`
	Performances []*struct {
		PlayID   string `json:"playID"`
		Audience int    `json:"audience"`
	} `json:"performances"`
}

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
