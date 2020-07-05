package companyroute

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/server/viewmodel"
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

	companyID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	result, getErr := s.companyService.GetCompanyByID(companyID)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleGetCompanyUserRating to handle a get company request
func (s *Controller) handleGetCompanyUserRating(c *gin.Context) {

	companyID, errID := s.getIDParameter(c.Param("company_id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	userID, errID := s.getIDParameter(c.Param("user_id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	if userID == 0 || companyID == 0 {
		restErr := resterrors.NewBadRequestError("The params company_id and user_id need to be sent")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	result, getErr := s.companyService.GetCompanyUserRating(companyID, userID)
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
		log.Println(err)
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
		log.Println(err)
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

// handleUpdateCompanyRating to handle a update company request
func (s *Controller) handleUpdateCompanyRating(c *gin.Context) {
	var companyRating entity.CompanyRating

	err := c.ShouldBindJSON(&companyRating)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	resCompanyRating, updateErr := s.companyService.UpdateCompanyRating(companyRating)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}

	c.JSON(http.StatusOK, companyStructToViewmodelResponse(*resCompanyRating))
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
		return 0, resterrors.NewBadRequestError("Param id should be a number")
	}

	return companyID, nil
}

func companyStructToViewmodelResponse(companyRating entity.CompanyRating) (vmCompanyRating viewmodel.CompanyRating) {

	vmCompanyRating.UserID = companyRating.UserID
	vmCompanyRating.CompanyID = companyRating.CompanyID
	vmCompanyRating.CustomerService = companyRating.CustomerService
	vmCompanyRating.CompanyClean = companyRating.CompanyClean
	vmCompanyRating.IceBeer = companyRating.IceBeer
	vmCompanyRating.GoodFood = companyRating.GoodFood
	vmCompanyRating.WouldGoBack = companyRating.WouldGoBack

	return vmCompanyRating
}
