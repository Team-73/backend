package mysqlutils

import (
	"fmt"
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows       = "no rows in result set"
	duplicatedKeyCode = 1062
)

// HandleMySQLError - handle mysql errors
func HandleMySQLError(errCode string, err error) *resterrors.RestErr {

	//if exists the error on mysql.MySQLError
	sqlErr, exists := err.(*mysql.MySQLError)
	if !exists {
		if strings.Contains(err.Error(), errorNoRows) {
			if errCode == "Error 0011: " {
				return resterrors.NewNotFoundError(fmt.Sprintf("%sInvalid user credentials", errCode))
			}
			return resterrors.NewNotFoundError(fmt.Sprintf("%sNo records find with the parameters", errCode))
		}
		return resterrors.NewInternalServerError(
			fmt.Sprintf("%sError database response: %s", errCode, err.Error()))
	}

	switch sqlErr.Number {
	case duplicatedKeyCode:
		duplicatedKey := between(sqlErr.Message, "key '", "_UNIQUE")
		duplicatedKeyValue := between(sqlErr.Message, "entry '", "' for key")
		return resterrors.NewBadRequestError(fmt.Sprintf("%sThe %s %s already exists", errCode, duplicatedKey, duplicatedKeyValue))
	}

	return resterrors.NewInternalServerError(
		fmt.Sprintf("%sError trying to processing database request: %s", errCode, err.Error()))
}

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:]
}
