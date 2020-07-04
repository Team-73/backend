package entity

// Company entity
type Company struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	CountryCode    string  `json:"country_code"`
	AreaCode       string  `json:"area_code"`
	PhoneNumber    string  `json:"phone_number"`
	DocumentNumber string  `json:"document_number"`
	Website        string  `json:"website"`
	BusinessID     string  `json:"business_id"`
	Address        Address `json:"address"`
}

// Address entity
type Address struct {
	Country        string `json:"country"`
	Street         string `json:"street"`
	Number         string `json:"number"`
	Complement     string `json:"complement"`
	ZipCode        string `json:"zip_code"`
	Neighborhood   string `json:"neighborhood"`
	City           string `json:"city"`
	FederativeUnit string `json:"federative_unit"`
}
