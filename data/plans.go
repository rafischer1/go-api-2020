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
var Plans = []Plan{{ID: 1, Name: "HMO NY-2500", Rates: Rates{EE: 200, ES: 250, EC: 300, FAM: 450}},
	{ID: 2, Name: "PPO CO-5000-XYZ", Rates: Rates{EE: 100, ES: 150, EC: 200, FAM: 350}},
	{ID: 3, Name: "HDHP Gold XPDS", Rates: Rates{EE: 350, ES: 450, EC: 500, FAM: 600}}}
