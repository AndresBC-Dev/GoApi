package models

type smartphone struct {
	Id            uint64
	Name          string
	Price         int
	CountryOrigin string
	OS            string
}

type CreateSmartphoneCMD struct {
	Name           string `json:"name"`
	Price          int    `jason:"price"`
	CountryOriginn string `json:"country_origin"`
	OS             string `json:"os"`
}
