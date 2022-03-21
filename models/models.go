package models

type Sklad struct {
	Id      int
	Name    string
	Firm    string
	Qtym    int
	Qtyt    int
	Price   string
	Cellm   string
	Cellt   string
	Partnum string
}

type Analog struct {
	Firm   string
	Number string
}

var Analogs []Analog
