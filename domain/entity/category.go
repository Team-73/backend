package entity

import (
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
)

// Category entity
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Validate to validate a user data
func (category *Category) Validate() *resterrors.RestErr {

	category.Name = strings.TrimSpace(strings.ToLower(category.Name))
	if category.Name == "" {
		return resterrors.NewBadRequestError("Name is invalid")
	}

	return nil
}
