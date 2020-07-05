package service

import (
	"strings"

	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type companyService struct {
	svc *Service
}

//newCompanyService return a new instance of the service
func newCompanyService(svc *Service) contract.CompanyService {
	return &companyService{
		svc: svc,
	}
}

func (s *companyService) GetCompanies() (*[]entity.Company, *resterrors.RestErr) {

	companies, err := s.svc.db.Company().GetCompanies()
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (s *companyService) GetCompanyByID(companyID int64) (*entity.Company, *resterrors.RestErr) {
	company := &entity.Company{
		ID: companyID,
	}

	company, err := s.svc.db.Company().GetCompanyByID(companyID)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *companyService) CreateCompany(company entity.Company) (int64, *resterrors.RestErr) {
	if err := company.Validate(); err != nil {
		return 0, err
	}

	newCompany, err := s.svc.db.Company().Create(company)
	if err != nil {
		return 0, err
	}

	return newCompany, nil
}

func (s *companyService) UpdateCompany(company entity.Company) (*entity.Company, *resterrors.RestErr) {

	currentCompany, err := s.GetCompanyByID(company.ID)
	if err != nil {
		return nil, err
	}

	if company.Name != "" {
		currentCompany.Name = strings.TrimSpace(company.Name)
	}

	if company.Email != "" {
		currentCompany.Email = strings.TrimSpace(company.Email)
	}

	if company.CountryCode != "" {
		currentCompany.CountryCode = strings.TrimSpace(company.CountryCode)
	}

	if company.AreaCode != "" {
		currentCompany.AreaCode = strings.TrimSpace(company.AreaCode)
	}

	if company.PhoneNumber != "" {
		currentCompany.PhoneNumber = strings.TrimSpace(company.PhoneNumber)
	}

	if company.BusinessID != 0 {
		currentCompany.BusinessID = company.BusinessID
	}

	updatedCompany, err := s.svc.db.Company().Update(*currentCompany)
	if err != nil {
		return nil, err
	}

	return updatedCompany, nil
}

func (s *companyService) DeleteCompany(companyID int64) *resterrors.RestErr {
	return s.svc.db.Company().Delete(companyID)
}
