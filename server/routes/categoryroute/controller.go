package categoryroute

import (
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

//Controller holds category handlers
type Controller struct {
	categoryService contract.CategoryService
}

//NewController to handle requests
func NewController(categoryService contract.CategoryService) *Controller {
	once.Do(func() {
		instance = &Controller{
			categoryService: categoryService,
		}
	})
	return instance
}

// handleGetCategorys to handle a get categories request
func (s *Controller) handleGetCategories(c *gin.Context) {

	result, getErr := s.categoryService.GetCategories()
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

// handleGetCategoryByID to handle a get category request
func (s *Controller) handleGetCategoryByID(c *gin.Context) {

	categoryID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	category, getErr := s.categoryService.GetCategoryByID(categoryID)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, category)
}

// handleCreateCategory to handle a create category request
func (s *Controller) handleCreateCategory(c *gin.Context) {

	var category entity.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	categoryID, createErr := s.categoryService.CreateCategory(category)
	if createErr != nil {
		c.JSON(createErr.StatusCode, createErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": categoryID})
}

// handleUpdateCategory to handle a update category request
func (s *Controller) handleUpdateCategory(c *gin.Context) {
	var category entity.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	categoryID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	category.ID = categoryID

	resCategory, updateErr := s.categoryService.UpdateCategory(category)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}

	c.JSON(http.StatusOK, resCategory)
}

// handleDeleteCategory to handle a delete category request
func (s *Controller) handleDeleteCategory(c *gin.Context) {

	categoryID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	deleteErr := s.categoryService.DeleteCategory(categoryID)
	if deleteErr != nil {
		c.JSON(deleteErr.StatusCode, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

func (s *Controller) getIDParameter(categoryParamID string) (int64, *resterrors.RestErr) {
	categoryID, categoryErr := strconv.ParseInt(categoryParamID, 10, 64)
	if categoryErr != nil {
		return 0, resterrors.NewBadRequestError("Category id should be a number")
	}

	return categoryID, nil
}
