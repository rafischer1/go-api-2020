package data

// TrekBook struct
type TrekBook struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Series     string         `json:"series"`
	Year       int            `json:"year"`
	BookReview TrekBookReview `json:"book_review"`
}

// TrekBookReview struct
type TrekBookReview struct {
	TrekBookID int    `json:"trek_book_id"`
	Author     string `json:"author"`
	Review     string `json:"year"`
}

// TrekBooks array
var TrekBooks = []TrekBook{{ID: 1, Name: "Spock's World", Year: 1993, Series: "TOS", BookReview: TrekBookReview{TrekBookID: 1, Author: "RAF", Review: "Awesome."}}, {ID: 2, Name: "Trek To Madworld", Year: 1979, Series: "TOS", BookReview: TrekBookReview{TrekBookID: 2, Author: "RAF", Review: "Meh."}}, {ID: 3, Name: "Gulliver's Fugitives", Year: 2001, Series: "TNG", BookReview: TrekBookReview{TrekBookID: 3, Author: "RAF", Review: "Yeah."}}, {ID: 4, Name: "Millennium", Year: 2003, Series: "DS9"}}
