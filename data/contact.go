package data

type Portfolio struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Github   string `json:"github"`
	LinkedIn string `json:"linkedIn"`
	Image    string `json:"imageUrl"`
}

var Personal = Portfolio{
	Name:     "Artie Fischer",
	Email:    "artiefischer@gmail.com",
	Github:   "https://github.com/rafischer1",
	LinkedIn: "https://www.linkedin.com/in/robert-a-fischer/",
	Image:    "https://avatars2.githubusercontent.com/u/39342327?s=460&v=4",
}

var WelcomeMessage = "Hi there! You made it to the backend - go to /contact for my info"
