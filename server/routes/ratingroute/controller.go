package ratingroute

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

//Controller holds rating handlers
type Controller struct {
	ratingService contract.RatingService
}

//NewController to handle requests
func NewController(ratingService contract.RatingService) *Controller {
	once.Do(func() {
		instance = &Controller{
			ratingService: ratingService,
		}
	})
	return instance
}

// handleGetCompanyUserRating to handle a get rating request
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

	result, getErr := s.ratingService.GetCompanyUserRating(companyID, userID)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleUpdateRating to handle a update rating request
func (s *Controller) handleUpdateRating(c *gin.Context) {
	var rating entity.Rating

	err := c.ShouldBindJSON(&rating)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	resRating, updateErr := s.ratingService.UpdateRating(rating)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}

	c.JSON(http.StatusOK, ratingStructToViewmodelResponse(*resRating))
}

func (s *Controller) getIDParameter(companyParamID string) (int64, *resterrors.RestErr) {
	companyID, companyErr := strconv.ParseInt(companyParamID, 10, 64)
	if companyErr != nil {
		return 0, resterrors.NewBadRequestError("Param id should be a number")
	}

	return companyID, nil
}

func ratingStructToViewmodelResponse(rating entity.Rating) (vmRating viewmodel.Rating) {

	vmRating.UserID = rating.UserID
	vmRating.CompanyID = rating.CompanyID
	vmRating.TotalRating = rating.TotalRating
	vmRating.CustomerService = rating.CustomerService
	vmRating.CompanyClean = rating.CompanyClean
	vmRating.IceBeer = rating.IceBeer
	vmRating.GoodFood = rating.GoodFood
	vmRating.WouldGoBack = rating.WouldGoBack

	return vmRating
}
