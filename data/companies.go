package data

// Company struct
type Company struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address Location `json:"address"`
}

// Location struct
type Location struct {
	Street  string `json:"street"`
	State   string `json:"state"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Country string `json:"country"`
}

// Companies arr
var Companies = []Company{{ID: 1, Name: "Artie's Fullstack Inc", Address: Location{
	Street:  "314 Main Street",
	State:   "CO",
	City:    "Denver",
	Zipcode: "80344",
	Country: "USA",
}}, {ID: 2, Name: "Leah's Content LLC", Address: Location{
	Street:  "3217 18th Street",
	City:    "Boulder",
	State:   "CO",
	Zipcode: "80304",
	Country: "USA",
}}, {ID: 3, Name: "Helen's Baby Toys", Address: Location{
	Street:  "1800 Congress Street",
	City:    "Portland",
	State:   "ME",
	Zipcode: "08301",
	Country: "USA",
}}, {ID: 3, Name: "Stevie's Dog Toys", Address: Location{
	Street:  "4 Bark Avenue",
	City:    "Philadelphia",
	State:   "PA",
	Zipcode: "17564",
	Country: "USA",
}}}
