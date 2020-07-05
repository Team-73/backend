package service

import (
	"strings"

	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type businessService struct {
	svc *Service
}

//newBusinessService return a new instance of the service
func newBusinessService(svc *Service) contract.BusinessService {
	return &businessService{
		svc: svc,
	}
}

func (s *businessService) GetBusinesses() (*[]entity.Business, *resterrors.RestErr) {

	businesss, err := s.svc.db.Business().GetBusinesses()
	if err != nil {
		return nil, err
	}

	return businesss, nil
}

func (s *businessService) GetBusinessByID(businessID int64) (*entity.Business, *resterrors.RestErr) {
	business := &entity.Business{
		ID: businessID,
	}

	business, err := s.svc.db.Business().GetBusinessByID(businessID)
	if err != nil {
		return nil, err
	}

	return business, nil
}

func (s *businessService) CreateBusiness(business entity.Business) (int64, *resterrors.RestErr) {
	if err := business.Validate(); err != nil {
		return 0, err
	}

	newBusinessID, err := s.svc.db.Business().Create(business)
	if err != nil {
		return 0, err
	}

	return newBusinessID, nil
}

func (s *businessService) UpdateBusiness(business entity.Business) (*entity.Business, *resterrors.RestErr) {

	currentBusiness, err := s.GetBusinessByID(business.ID)
	if err != nil {
		return nil, err
	}

	if business.Name != "" {
		currentBusiness.Name = strings.TrimSpace(business.Name)
	}

	updatedBusiness, err := s.svc.db.Business().Update(*currentBusiness)
	if err != nil {
		return nil, err
	}

	return updatedBusiness, nil
}

func (s *businessService) DeleteBusiness(businessID int64) *resterrors.RestErr {
	return s.svc.db.Business().Delete(businessID)
}
