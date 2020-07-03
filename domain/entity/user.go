package entity

import (
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
)

// User entity
type User struct {
	ID             int64   `json:"id,omitempty"`
	Name           string  `json:"name,omitempty"`
	Email          string  `json:"email,omitempty"`
	Password       string  `json:"password,omitempty"`
	DocumentNumber string  `json:"document_number,omitempty"`
	AreaCode       string  `json:"area_code,omitempty"`
	PhoneNumber    string  `json:"phone_number,omitempty"`
	Birthdate      string  `json:"birthdate,omitempty"`
	Gender         string  `json:"gender,omitempty"`
	Revenue        float64 `json:"revenue,omitempty"` //this information is the "frenda" field on Gr1d API
	Active         bool    `json:"active,omitempty"`
}

// Validate to validate a user data
func (user *User) Validate() *resterrors.RestErr {

	user.Name = strings.TrimSpace(user.Name)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return resterrors.NewBadRequestError("Invalid email address")
	}

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
