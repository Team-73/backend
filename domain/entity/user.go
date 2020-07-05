package entity

import (
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
)

// User entity
type User struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	DocumentNumber string  `json:"document_number"`
	CountryCode    string  `json:"country_code"`
	AreaCode       string  `json:"area_code"`
	PhoneNumber    string  `json:"phone_number"`
	Birthdate      string  `json:"birthdate"`
	Gender         string  `json:"gender"`
	Revenue        float64 `json:"revenue"` //this information is the "frenda" field on Gr1d API
	Active         bool    `json:"active"`
}

// Validate to validate a user data
func (user *User) Validate() *resterrors.RestErr {

	user.Birthdate = strings.TrimSpace(user.Name)
	if user.Birthdate == "" {
		user.Birthdate = "1993-10-25"
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return resterrors.NewBadRequestError("Email address is invalid")
	}

	/* user.CountryCode = strings.TrimSpace(user.CountryCode)
	if user.CountryCode == "" {
		return resterrors.NewBadRequestError("Country code is invalid")
	}

	user.AreaCode = strings.TrimSpace(user.AreaCode)
	if user.AreaCode == "" {
		return resterrors.NewBadRequestError("Area code is invalid")
	}

	user.PhoneNumber = strings.TrimSpace(user.PhoneNumber)
	if user.PhoneNumber == "" {
		return resterrors.NewBadRequestError("Phone number is invalid")
	} */

	user.Password = strings.TrimSpace(user.Password)
	err := user.validadePassword()
	if err != nil {
		return err
	}

	return nil
}

func (user *User) validadePassword() *resterrors.RestErr {
	if user.Password == "" || len(user.Password) < 8 {
		return resterrors.NewBadRequestError("Password need at least 8 caracters")
	}

	return nil
}

//LoginRequest entity
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
