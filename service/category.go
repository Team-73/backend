package service

import (
	"strings"

	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type categoryService struct {
	svc *Service
}

//newCategoryService return a new instance of the service
func newCategoryService(svc *Service) contract.CategoryService {
	return &categoryService{
		svc: svc,
	}
}

func (s *categoryService) GetCategories() (*[]entity.Category, *resterrors.RestErr) {

	categories, err := s.svc.db.Category().GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *categoryService) GetCategoryByID(categoryID int64) (*entity.Category, *resterrors.RestErr) {
	category := &entity.Category{
		ID: categoryID,
	}

	category, err := s.svc.db.Category().GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) CreateCategory(category entity.Category) (int64, *resterrors.RestErr) {
	if err := category.Validate(); err != nil {
		return 0, err
	}

	newCategoryID, err := s.svc.db.Category().Create(category)
	if err != nil {
		return 0, err
	}

	return newCategoryID, nil
}

func (s *categoryService) UpdateCategory(category entity.Category) (*entity.Category, *resterrors.RestErr) {

	currentCategory, err := s.GetCategoryByID(category.ID)
	if err != nil {
		return nil, err
	}

	if category.Name != "" {
		currentCategory.Name = strings.TrimSpace(category.Name)
	}

	updatedCategory, err := s.svc.db.Category().Update(*currentCategory)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func (s *categoryService) DeleteCategory(categoryID int64) *resterrors.RestErr {
	return s.svc.db.Category().Delete(categoryID)
}
