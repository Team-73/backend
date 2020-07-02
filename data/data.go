package data

import (
	"github.com/Team-73/backend/data/mysql"
	"github.com/Team-73/backend/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return mysql.Instance()
}
