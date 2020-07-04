package productroute

import (
	"github.com/gin-gonic/gin"
)

// Route holds the product handlers
type Route struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRoute returns a new Route instance
func NewRoute(ctrl *Controller, router *gin.Engine) *Route {
	return &Route{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of user requests
func (r *Route) RegisterRoutes() {

	r.router.GET("/products", r.ctrl.handleGetProducts)
	r.router.GET("/product/:id", r.ctrl.handleGetProductByID)
	r.router.POST("/product", r.ctrl.handleCreateProduct)
	r.router.PUT("/product/:id", r.ctrl.handleUpdateProduct)
	r.router.DELETE("/product/:id", r.ctrl.handleDeleteProduct)

}
