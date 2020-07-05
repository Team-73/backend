package businessroute

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

//Controller holds business handlers
type Controller struct {
	businessService contract.BusinessService
}

//NewController to handle requests
func NewController(businessService contract.BusinessService) *Controller {
	once.Do(func() {
		instance = &Controller{
			businessService: businessService,
		}
	})
	return instance
}

// handleGetBusinesses to handle a get businesses request
func (s *Controller) handleGetBusinesses(c *gin.Context) {

	result, getErr := s.businessService.GetBusinesses()
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

// handleGetBusinessByID to handle a get business request
func (s *Controller) handleGetBusinessByID(c *gin.Context) {

	businessID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	business, getErr := s.businessService.GetBusinessByID(businessID)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, business)
}

// handleCreateBusiness to handle a create business request
func (s *Controller) handleCreateBusiness(c *gin.Context) {

	var business entity.Business

	err := c.ShouldBindJSON(&business)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	businessID, createErr := s.businessService.CreateBusiness(business)
	if createErr != nil {
		c.JSON(createErr.StatusCode, createErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": businessID})
}

// handleUpdateBusiness to handle a update business request
func (s *Controller) handleUpdateBusiness(c *gin.Context) {
	var business entity.Business

	err := c.ShouldBindJSON(&business)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	businessID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	business.ID = businessID

	resBusiness, updateErr := s.businessService.UpdateBusiness(business)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}

	c.JSON(http.StatusOK, resBusiness)
}

// handleDeleteBusiness to handle a delete business request
func (s *Controller) handleDeleteBusiness(c *gin.Context) {

	businessID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	deleteErr := s.businessService.DeleteBusiness(businessID)
	if deleteErr != nil {
		c.JSON(deleteErr.StatusCode, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

func (s *Controller) getIDParameter(businessParamID string) (int64, *resterrors.RestErr) {
	businessID, businessErr := strconv.ParseInt(businessParamID, 10, 64)
	if businessErr != nil {
		return 0, resterrors.NewBadRequestError("Param id should be a number")
	}

	return businessID, nil
}
