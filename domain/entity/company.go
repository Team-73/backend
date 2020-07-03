package entity

import "time"

// Company entity
type Company struct {
	ID             int64   `json:"id"`
	CompanyName    string  `json:"company_name,omitempty"`
	AreaCode       string  `json:"area_code,omitempty"`
	PhoneNumber    string  `json:"phone_number,omitempty"`
	Email          string  `json:"email,omitempty"`
	DocumentNumber string  `json:"document_number,omitempty"`
	CommercialName string  `json:"commercial_name,omitempty"`
	Website        string  `json:"website,omitempty"`
	Business       string  `json:"business,omitempty"`
	Address        Address `json:"address,omitempty"`
	Partner        Partner `json:"partner,omitempty"`
}

// CreateCompany entity
type CreateCompany struct {
	CompanyName    string  `json:"company_name,omitempty"`
	AreaCode       string  `json:"area_code,omitempty"`
	PhoneNumber    string  `json:"phone_number,omitempty"`
	Email          string  `json:"email,omitempty"`
	DocumentNumber string  `json:"document_number,omitempty"`
	CommercialName string  `json:"commercial_name,omitempty"`
	Website        string  `json:"website,omitempty"`
	Business       string  `json:"business,omitempty"`
	Address        Address `json:"address,omitempty"`
	Partner        Partner `json:"partner,omitempty"`
}

// Address entity
type Address struct {
	Street         string `json:"street,omitempty"`
	Number         string `json:"number,omitempty"`
	Complement     string `json:"complement,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
	Neighborhood   string `json:"neighborhood,omitempty"`
	City           string `json:"city,omitempty"`
	FederativeUnit string `json:"federative_unit,omitempty"`
}

// Partner entity
type Partner struct {
	Name           string    `json:"name,omitempty"`
	DocumentNumber string    `json:"document_number,omitempty"`
	Email          string    `json:"email,omitempty"`
	AreaCode       string    `json:"area_code,omitempty"`
	PhoneNumber    string    `json:"phone_number,omitempty"`
	Birthdate      time.Time `json:"birthdate,omitempty"`
	Gender         string    `json:"gender,omitempty"`
}
