package entity

import "time"

// Company entity
type Company struct {
	ID             int64   `json:"id"`
	CompanyName    string  `json:"company_name"`
	AreaCode       string  `json:"area_code"`
	PhoneNumber    string  `json:"phone_number"`
	Email          string  `json:"email"`
	DocumentNumber string  `json:"document_number"`
	CommercialName string  `json:"commercial_name"`
	Website        string  `json:"website"`
	Business       string  `json:"business"`
	Address        Address `json:"address"`
	Partner        Partner `json:"partner"`
}

// CreateCompany entity
type CreateCompany struct {
	CompanyName    string  `json:"company_name"`
	AreaCode       string  `json:"area_code"`
	PhoneNumber    string  `json:"phone_number"`
	Email          string  `json:"email"`
	DocumentNumber string  `json:"document_number"`
	CommercialName string  `json:"commercial_name"`
	Website        string  `json:"website"`
	Business       string  `json:"business"`
	Address        Address `json:"address"`
	Partner        Partner `json:"partner"`
}

// Address entity
type Address struct {
	Street         string `json:"street"`
	Number         string `json:"number"`
	Complement     string `json:"complement"`
	ZipCode        string `json:"zip_code"`
	Neighborhood   string `json:"neighborhood"`
	City           string `json:"city"`
	FederativeUnit string `json:"federative_unit"`
}

// Partner entity
type Partner struct {
	Name           string    `json:"name"`
	DocumentNumber string    `json:"document_number"`
	Email          string    `json:"email"`
	AreaCode       string    `json:"area_code"`
	PhoneNumber    string    `json:"phone_number"`
	Birthdate      time.Time `json:"birthdate"`
	Gender         string    `json:"gender"`
}
