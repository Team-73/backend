package service

import "github.com/Team-73/backend/domain/contract"

// Service holds the domain service repositories
type Service struct {
	db contract.RepoManager
}

// New returns a new domain Service instance
func New(db contract.RepoManager) *Service {
	svc := new(Service)
	svc.db = db

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	BusinessService(svc *Service) contract.BusinessService
	CategoryService(svc *Service) contract.CategoryService
	CompanyService(svc *Service) contract.CompanyService
	ProductService(svc *Service) contract.ProductService
	UserService(svc *Service) contract.UserService
	RatingService(svc *Service) contract.RatingService
	OrderService(svc *Service) contract.OrderService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) BusinessService(svc *Service) contract.BusinessService {
	return newBusinessService(svc)
}

func (s *serviceManager) CategoryService(svc *Service) contract.CategoryService {
	return newCategoryService(svc)
}

func (s *serviceManager) CompanyService(svc *Service) contract.CompanyService {
	return newCompanyService(svc)
}

func (s *serviceManager) ProductService(svc *Service) contract.ProductService {
	return newProductService(svc)
}

func (s *serviceManager) UserService(svc *Service) contract.UserService {
	return newUserService(svc)
}

func (s *serviceManager) RatingService(svc *Service) contract.RatingService {
	return newRatingService(svc)
}

func (s *serviceManager) OrderService(svc *Service) contract.OrderService {
	return newOrderService(svc)
}
