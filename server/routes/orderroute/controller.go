package orderroute

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

//Controller holds user handlers
type Controller struct {
	orderService contract.OrderService
}

//NewController to handle requests
func NewController(orderService contract.OrderService) *Controller {
	once.Do(func() {
		instance = &Controller{
			orderService: orderService,
		}
	})
	return instance
}

// handleCreateOrder - handle a order request
func (s *Controller) handleCreateOrder(c *gin.Context) {

	var order entity.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	orderID, createErr := s.orderService.CreateOrder(order)
	if createErr != nil {
		c.JSON(createErr.StatusCode, createErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": orderID})

}

// handleGetOrderByUserID - handle a get order request
func (s *Controller) handleGetOrderByUserID(c *gin.Context) {

	userID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	orders, err := s.orderService.GetOrderByUserID(userID)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, orders)

}

func (s *Controller) getIDParameter(ParamID string) (int64, *resterrors.RestErr) {
	id, Err := strconv.ParseInt(ParamID, 10, 64)
	if Err != nil {
		return 0, resterrors.NewBadRequestError("Param id should be a number")
	}

	return id, nil
}
