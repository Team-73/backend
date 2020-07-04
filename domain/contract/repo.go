package contract

import (
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Ping() PingRepo
	User() UserRepo
	Product() ProductRepo
	Category() CategoryRepo
}

// PingRepo defines the data set for a ping
type PingRepo interface{}

// UserRepo defines the data set for a user
type UserRepo interface {
	GetUsers() (*[]entity.User, *resterrors.RestErr)
	GetUserByID(userID int64) (*entity.User, *resterrors.RestErr)
	Create(entity.User) (int64, *resterrors.RestErr)
	Update(entity.User) (*entity.User, *resterrors.RestErr)
	Delete(userID int64) *resterrors.RestErr
	GetByEmailAndPassword(user entity.LoginRequest) (*entity.User, *resterrors.RestErr)
}

// ProductRepo defines the data set for a product
type ProductRepo interface {
	GetProducts() (*[]entity.Product, *resterrors.RestErr)
	GetProductByID(productID int64) (*entity.Product, *resterrors.RestErr)
	Create(entity.Product) (int64, *resterrors.RestErr)
	Update(entity.Product) (*entity.Product, *resterrors.RestErr)
	Delete(productID int64) *resterrors.RestErr
}

// CategoryRepo defines the data set for a category
type CategoryRepo interface {
	GetCategories() (*[]entity.Category, *resterrors.RestErr)
	GetCategoryByID(categoryID int64) (*entity.Category, *resterrors.RestErr)
	Create(entity.Category) (int64, *resterrors.RestErr)
	Update(entity.Category) (*entity.Category, *resterrors.RestErr)
	Delete(categoryID int64) *resterrors.RestErr
}
