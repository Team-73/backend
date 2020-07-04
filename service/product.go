package service

import (
	"strings"

	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type productService struct {
	svc *Service
}

//newProductService return a new instance of the service
func newProductService(svc *Service) contract.ProductService {
	return &productService{
		svc: svc,
	}
}

func (s *productService) GetProducts(categoryID int64) (*[]entity.Product, *resterrors.RestErr) {

	products, err := s.svc.db.Product().GetProducts(categoryID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) GetProductByID(productID int64) (*entity.Product, *resterrors.RestErr) {
	product := &entity.Product{
		ID: productID,
	}

	product, err := s.svc.db.Product().GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) CreateProduct(product entity.Product) (int64, *resterrors.RestErr) {
	if err := product.Validate(); err != nil {
		return 0, err
	}

	newProductID, err := s.svc.db.Product().Create(product)
	if err != nil {
		return 0, err
	}

	return newProductID, nil
}

func (s *productService) UpdateProduct(product entity.Product) (*entity.Product, *resterrors.RestErr) {

	currentProduct, err := s.GetProductByID(product.ID)
	if err != nil {
		return nil, err
	}

	if product.Name != "" {
		currentProduct.Name = strings.TrimSpace(product.Name)
	}

	if product.Description != "" {
		currentProduct.Description = strings.TrimSpace(product.Description)
	}

	if product.Price != 0 {
		currentProduct.Price = product.Price
	}

	if product.DiscountPrice != 0 {
		currentProduct.DiscountPrice = product.DiscountPrice
	}

	if product.CategoryID != 0 {
		currentProduct.CategoryID = product.CategoryID
	}

	if product.MinimumAgeForConsumption != 0 {
		currentProduct.MinimumAgeForConsumption = product.MinimumAgeForConsumption
	}

	if product.ProductImageURL != "" {
		currentProduct.ProductImageURL = strings.TrimSpace(product.ProductImageURL)
	}

	if product.TimeForPreparingMinutes != 0 {
		currentProduct.TimeForPreparingMinutes = product.TimeForPreparingMinutes
	}

	updatedProduct, err := s.svc.db.Product().Update(*currentProduct)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (s *productService) DeleteProduct(productID int64) *resterrors.RestErr {
	return s.svc.db.Product().Delete(productID)
}
