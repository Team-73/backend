package productroute

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/gin-gonic/gin"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds product handlers
type Controller struct {
	productService contract.ProductService
}

//NewController to handle requests
func NewController(productService contract.ProductService) *Controller {
	once.Do(func() {
		instance = &Controller{
			productService: productService,
		}
	})
	return instance
}

// handleGetProducts to handle a get product request
func (s *Controller) handleGetProducts(c *gin.Context) {

	var categoryID int64 = 0
	var err error

	category := c.Query("category")
	if category != "" {

		categoryID, err = strconv.ParseInt(c.Query("category"), 10, 64)
		if err != nil {
			log.Println(err)
			restErr := resterrors.NewBadRequestError("Unable to parse category value")
			c.JSON(restErr.StatusCode, restErr)
			return
		}
	}

	result, getErr := s.productService.GetProducts(categoryID)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	if len(*result) == 0 {
		notFound := resterrors.NewNotFoundError("No records find with the parameters")
		c.JSON(http.StatusOK, notFound)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleGetProductByID to handle a get product request
func (s *Controller) handleGetProductByID(c *gin.Context) {

	productID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	product, getErr := s.productService.GetProductByID(productID)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, product)
}

// handleCreateProduct to handle a create product request
func (s *Controller) handleCreateProduct(c *gin.Context) {

	var product entity.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	productID, createErr := s.productService.CreateProduct(product)
	if createErr != nil {
		c.JSON(createErr.StatusCode, createErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": productID})
}

// handleUpdateProduct to handle a update product request
func (s *Controller) handleUpdateProduct(c *gin.Context) {
	var product entity.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	productID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	product.ID = productID

	resProduct, updateErr := s.productService.UpdateProduct(product)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}

	c.JSON(http.StatusOK, resProduct)
}

// handleDeleteProduct to handle a delete product request
func (s *Controller) handleDeleteProduct(c *gin.Context) {

	productID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	deleteErr := s.productService.DeleteProduct(productID)
	if deleteErr != nil {
		c.JSON(deleteErr.StatusCode, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

func (s *Controller) getIDParameter(productParamID string) (int64, *resterrors.RestErr) {
	productID, productErr := strconv.ParseInt(productParamID, 10, 64)
	if productErr != nil {
		return 0, resterrors.NewBadRequestError("Product id should be a number")
	}

	return productID, nil
}
