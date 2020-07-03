package contract

import (
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

// PingService holds a ping service operations
type PingService interface {
}

// UserService holds a user service operations
type UserService interface {
	GetUser(userID int64) (*entity.User, *resterrors.RestErr)
	CreateUser(entity.User) (int64, *resterrors.RestErr)
	UpdateUser(entity.User) (*entity.User, *resterrors.RestErr)
	DeleteUser(userID int64) *resterrors.RestErr
	LoginUser(request entity.LoginRequest) (*entity.User, *resterrors.RestErr)
}
