package data

// Portfolio struct defined
type Portfolio struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Github   string `json:"Github"`
	LinkedIn string `json:"LinkedIn"`
	Resume   string `json:"Resume Download"`
}

// Personal contact information
var Personal = Portfolio{
	Name:     "Artie Fischer",
	Email:    "artiefischer@gmail.com",
	Github:   "https://github.com/rafischer1",
	LinkedIn: "https://www.linkedin.com/in/robert-a-fischer/",
	Resume:   "http://localhost:8080/Robert_Fischer_Resume_2020",
}

// WelcomeMessage to pass to html template
var WelcomeMessage = "Hi there! You made it to the backend - go to /contact for my info"
