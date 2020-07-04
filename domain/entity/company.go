package entity

import (
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
)

// Company entity
type Company struct {
	ID             int64         `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	Email          string        `json:"email,omitempty"`
	CountryCode    string        `json:"country_code,omitempty"`
	AreaCode       string        `json:"area_code,omitempty"`
	PhoneNumber    string        `json:"phone_number,omitempty"`
	DocumentNumber string        `json:"document_number,omitempty"`
	Website        string        `json:"website,omitempty"`
	BusinessID     int64         `json:"business_id,omitempty"`
	Address        Address       `json:"address,omitempty"`
	SocialNetwork  SocialNetwork `json:"social_network,omitempty"`
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

// SocialNetwork entity
type SocialNetwork struct {
	InstagramURL string `json:"instagram_url,omitempty"`
	FacebookURL  string `json:"facebook_url,omitempty"`
	LinkedinURL  string `json:"linkedin_url,omitempty"`
	TwitterURL   string `json:"twitter_url,omitempty"`
}

// Validate to validate a user data
func (company *Company) Validate() *resterrors.RestErr {

	company.Name = strings.TrimSpace(company.Name)
	if company.Name == "" {
		return resterrors.NewBadRequestError("Name is invalid")
	}

	company.Email = strings.TrimSpace(company.Email)
	if company.Email != "" {
		return resterrors.NewBadRequestError("Email is invalid")
	}

	company.CountryCode = strings.TrimSpace(company.CountryCode)
	if company.CountryCode != "" {
		return resterrors.NewBadRequestError("CountryCode is invalid")
	}

	company.AreaCode = strings.TrimSpace(company.AreaCode)
	if company.AreaCode != "" {
		return resterrors.NewBadRequestError("AreaCode is invalid")
	}

	company.PhoneNumber = strings.TrimSpace(company.PhoneNumber)
	if company.PhoneNumber != "" {
		return resterrors.NewBadRequestError("PhoneNumber is invalid")
	}

	if company.BusinessID != 0 {
		return resterrors.NewBadRequestError("BusinessID is invalid")
	}

	return nil
}
