package entity

import (
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
)

// Business entity
type Business struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Validate to validate a user data
func (business *Business) Validate() *resterrors.RestErr {

	business.Name = strings.TrimSpace(business.Name)
	if business.Name == "" {
		return resterrors.NewBadRequestError("Name is invalid")
	}

	return nil
}
