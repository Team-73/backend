package viewmodel

// User viewModel
type User struct {
	ID             int64   `json:"id,omitempty"`
	Name           string  `json:"name,omitempty"`
	Email          string  `json:"email,omitempty"`
	DocumentNumber string  `json:"document_number,omitempty"`
	AreaCode       string  `json:"area_code,omitempty"`
	PhoneNumber    string  `json:"phone_number,omitempty"`
	Birthdate      string  `json:"birthdate,omitempty"`
	Gender         string  `json:"gender,omitempty"`
	Revenue        float64 `json:"revenue,omitempty"` //this information is the "frenda" field on Gr1d API
	Active         bool    `json:"active,omitempty"`
}
