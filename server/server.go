package server

import (
	"github.com/Team-73/backend/server/routes/businessroute"
	"github.com/Team-73/backend/server/routes/categoryroute"
	"github.com/Team-73/backend/server/routes/companyroute"
	"github.com/Team-73/backend/server/routes/orderroute"
	"github.com/Team-73/backend/server/routes/pingroute"
	"github.com/Team-73/backend/server/routes/productroute"
	"github.com/Team-73/backend/server/routes/ratingroute"
	"github.com/Team-73/backend/server/routes/userroute"
	"github.com/Team-73/backend/service"
	"github.com/gin-gonic/gin"
)

type controller struct {
	pingController     *pingroute.Controller
	userController     *userroute.Controller
	productController  *productroute.Controller
	categoryController *categoryroute.Controller
	businessController *businessroute.Controller
	companyController  *companyroute.Controller
	ratingController   *ratingroute.Controller
	orderController    *orderroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	svm := service.NewServiceManager()
	srv := gin.Default()

	return setupRoutes(srv, &controller{
		pingController:     pingroute.NewController(),
		userController:     userroute.NewController(svm.UserService(svc)),
		productController:  productroute.NewController(svm.ProductService(svc)),
		categoryController: categoryroute.NewController(svm.CategoryService(svc)),
		businessController: businessroute.NewController(svm.BusinessService(svc)),
		companyController:  companyroute.NewController(svm.CompanyService(svc)),
		ratingController:   ratingroute.NewController(svm.RatingService(svc)),
		orderController:    orderroute.NewController(svm.OrderService(svc)),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	pingroute.NewRoute(s.pingController, srv).RegisterRoutes()
	userroute.NewRoute(s.userController, srv).RegisterRoutes()
	productroute.NewRoute(s.productController, srv).RegisterRoutes()
	categoryroute.NewRoute(s.categoryController, srv).RegisterRoutes()
	businessroute.NewRoute(s.businessController, srv).RegisterRoutes()
	companyroute.NewRoute(s.companyController, srv).RegisterRoutes()
	ratingroute.NewRoute(s.ratingController, srv).RegisterRoutes()
	orderroute.NewRoute(s.orderController, srv).RegisterRoutes()

	return srv
}
