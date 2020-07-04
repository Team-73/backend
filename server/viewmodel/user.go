package viewmodel

// User viewModel
type User struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	DocumentNumber string  `json:"document_number"`
	CountryCode    string  `json:"country_code"`
	AreaCode       string  `json:"area_code"`
	PhoneNumber    string  `json:"phone_number"`
	Birthdate      string  `json:"birthdate"`
	Gender         string  `json:"gender"`
	Revenue        float64 `json:"revenue"` //this information is the "frenda" field on Gr1d API
	Active         bool    `json:"active"`
}
