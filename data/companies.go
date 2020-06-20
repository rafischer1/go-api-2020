package data

// Company struct for the the psql table subscribers
type Company struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address Location `json:"address"`
}

type Location struct {
	Street  string `json:"street"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
	Country string `json:"country"`
}

// Plans arr
var Companies = []Company{{ID: 1, Name: "Artie's Company 1", Address: Location{
	Street:  "hi",
	State:   "",
	Zipcode: "",
	Country: "",
}}}
