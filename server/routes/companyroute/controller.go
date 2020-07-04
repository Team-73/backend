package companyroute

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

//Controller holds company handlers
type Controller struct {
	companyService contract.CompanyService
}

//NewController to handle requests
func NewController(companyService contract.CompanyService) *Controller {
	once.Do(func() {
		instance = &Controller{
			companyService: companyService,
		}
	})
	return instance
}

// handleGetCompanyByID to handle a get company request
func (s *Controller) handleGetCompanyByID(c *gin.Context) {

	var companyID int64 = 0
	var err error

	company := c.Query("company")
	if company != "" {

		companyID, err = strconv.ParseInt(c.Query("company"), 10, 64)
		if err != nil {
			log.Println(err)
			restErr := resterrors.NewBadRequestError("Unable to parse company value")
			c.JSON(restErr.StatusCode, restErr)
			return
		}
	} else {
		companyID = 0
	}

	result, getErr := s.companyService.GetCompanyByID(companyID)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleGetCompanies to handle a get company request
func (s *Controller) handleGetCompanies(c *gin.Context) {

	result, getErr := s.companyService.GetCompanies()
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

// handleCreateCompany to handle a create company request
func (s *Controller) handleCreateCompany(c *gin.Context) {

	var company entity.Company

	err := c.ShouldBindJSON(&company)
	if err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	companyID, createErr := s.companyService.CreateCompany(company)
	if createErr != nil {
		c.JSON(createErr.StatusCode, createErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": companyID})
}

// handleUpdateCompany to handle a update company request
func (s *Controller) handleUpdateCompany(c *gin.Context) {
	var company entity.Company

	err := c.ShouldBindJSON(&company)
	if err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	companyID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	company.ID = companyID

	resCompany, updateErr := s.companyService.UpdateCompany(company)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}

	c.JSON(http.StatusOK, resCompany)
}

// handleDeleteCompany to handle a delete company request
func (s *Controller) handleDeleteCompany(c *gin.Context) {

	companyID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	deleteErr := s.companyService.DeleteCompany(companyID)
	if deleteErr != nil {
		c.JSON(deleteErr.StatusCode, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

func (s *Controller) getIDParameter(companyParamID string) (int64, *resterrors.RestErr) {
	companyID, companyErr := strconv.ParseInt(companyParamID, 10, 64)
	if companyErr != nil {
		return 0, resterrors.NewBadRequestError("Company id should be a number")
	}

	return companyID, nil
}
