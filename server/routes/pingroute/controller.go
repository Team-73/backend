package pingroute

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds user handlers
type Controller struct {
}

//NewController to handle requests
func NewController() *Controller {
	once.Do(func() {
		instance = &Controller{}
	})
	return instance
}

// handlePing - handle a Ping request
func (s *Controller) handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
