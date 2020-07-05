package contract

import (
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Ping() PingRepo
	Business() BusinessRepo
	Category() CategoryRepo
	Company() CompanyRepo
	Product() ProductRepo
	User() UserRepo
	Order() OrderRepo
	Rating() RatingRepo
}

// PingRepo defines the data set for a ping
type PingRepo interface{}

// BusinessRepo defines the data set for a business
type BusinessRepo interface {
	GetBusinesses() (*[]entity.Business, *resterrors.RestErr)
	GetBusinessByID(businessID int64) (*entity.Business, *resterrors.RestErr)
	Create(entity.Business) (int64, *resterrors.RestErr)
	Update(entity.Business) (*entity.Business, *resterrors.RestErr)
	Delete(businessID int64) *resterrors.RestErr
}

// CategoryRepo defines the data set for a category
type CategoryRepo interface {
	GetCategories() (*[]entity.Category, *resterrors.RestErr)
	GetCategoryByID(categoryID int64) (*entity.Category, *resterrors.RestErr)
	Create(entity.Category) (int64, *resterrors.RestErr)
	Update(entity.Category) (*entity.Category, *resterrors.RestErr)
	Delete(categoryID int64) *resterrors.RestErr
}

// CompanyRepo defines the data set for a company
type CompanyRepo interface {
	GetCompanies() ([]entity.Companies, *resterrors.RestErr)
	GetCompanyByID(companyID int64) (*entity.CompanyDetail, *resterrors.RestErr)
	Create(entity.CompanyDetail) (int64, *resterrors.RestErr)
	Update(entity.CompanyDetail) (*entity.CompanyDetail, *resterrors.RestErr)
	Delete(companyID int64) *resterrors.RestErr
}

// ProductRepo defines the data set for a product
type ProductRepo interface {
	GetProducts(categoryID int64) (*[]entity.Product, *resterrors.RestErr)
	GetProductByID(productID int64) (*entity.Product, *resterrors.RestErr)
	Create(entity.Product) (int64, *resterrors.RestErr)
	Update(entity.Product) (*entity.Product, *resterrors.RestErr)
	Delete(productID int64) *resterrors.RestErr
}

// UserRepo defines the data set for a user
type UserRepo interface {
	GetUsers() (*[]entity.User, *resterrors.RestErr)
	GetUserByID(userID int64) (*entity.User, *resterrors.RestErr)
	Create(entity.User) (int64, *resterrors.RestErr)
	Update(entity.User) (*entity.User, *resterrors.RestErr)
	Delete(userID int64) *resterrors.RestErr
	GetByEmailAndPassword(user entity.LoginRequest) (*entity.User, *resterrors.RestErr)
}

// OrderRepo defines the data set for a category
type OrderRepo interface {
	GetOrdersByUserID(userID int64) (*[]entity.OrdersByUserID, *resterrors.RestErr)
	GetOrderDetail(orderID int64) (*entity.OrderDetail, *resterrors.RestErr)
	GetOrderProducts(orderID int64) ([]entity.ProductsDetail, *resterrors.RestErr)
	CreateOrder(entity.Order) (int64, *resterrors.RestErr)
	CreateOrderProductAndReturnProductPrice(orderID int64, oderProduct entity.OrderProduct) (float64, *resterrors.RestErr)
	UpdateOrder(orderID int64, order entity.Order) *resterrors.RestErr
}

// RatingRepo defines the data set for a category
type RatingRepo interface {
	GetCompanyUserRating(companyID, userID int64) (*entity.Rating, *resterrors.RestErr)
	CreateRating(rating entity.Rating) (*entity.Rating, *resterrors.RestErr)
	UpdateRating(rating entity.Rating) (*entity.Rating, *resterrors.RestErr)
}
