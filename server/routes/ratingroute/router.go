package ratingroute

import (
	"github.com/gin-gonic/gin"
)

// Router holds the rating handlers
type Router struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRoute returns a new Route instance
func NewRoute(ctrl *Controller, router *gin.Engine) *Router {
	return &Router{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of rating requests
func (r *Router) RegisterRoutes() {

	r.router.GET("/rating/:company_id/:user_id", r.ctrl.handleGetCompanyUserRating)
	r.router.PUT("/rating", r.ctrl.handleUpdateRating)

}
