package data

// Plan struct for the the psql table subscribers
type Plan struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Rates Rates  `json:"rates"`
}

// Rates struct to display current cov type rates
type Rates struct {
	EE  int `json:"ee"`
	ES  int `json:"es"`
	EC  int `json:"ec"`
	FAM int `json:"fam"`
}

// Plans arr
var Plans = []Plan{{ID: 1, Name: "Artie's Plan 1", Rates: Rates{EE: 200, ES: 250, EC: 300, FAM: 450}}}
