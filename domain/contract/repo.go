package contract

import (
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Ping() PingRepo
	User() UserRepo
}

// PingRepo defines the data set for ping
type PingRepo interface{}

// UserRepo defines the data set for user
type UserRepo interface {
	GetByID(userID int64) (*entity.User, *resterrors.RestErr)
	Create(entity.User) (int64, *resterrors.RestErr)
	Update(entity.User) (*entity.User, *resterrors.RestErr)
	Delete(userID int64) *resterrors.RestErr
	GetByEmailAndPassword(user entity.LoginRequest) (*entity.User, *resterrors.RestErr)
}
