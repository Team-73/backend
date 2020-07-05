package userroute

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

//Controller holds user handlers
type Controller struct {
	userService contract.UserService
}

//NewController to handle requests
func NewController(userService contract.UserService) *Controller {
	once.Do(func() {
		instance = &Controller{
			userService: userService,
		}
	})
	return instance
}

// handleGetUserByID to handle a get user request
func (s *Controller) handleGetUserByID(c *gin.Context) {

	userID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	user, getErr := s.userService.GetUserByID(userID)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	vmResponse := userStructToViewmodelResponse(*user)

	c.JSON(http.StatusOK, vmResponse)
}

// handleGetUsers to handle a get user request
func (s *Controller) handleGetUsers(c *gin.Context) {

	result, getErr := s.userService.GetUsers()
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	if len(*result) == 0 {
		notFound := resterrors.NewNotFoundError("No records find with the parameters")
		c.JSON(http.StatusOK, notFound)
		return
	}

	vmResponse := userListStructToViewmodelResponse(*result)

	c.JSON(http.StatusOK, vmResponse)
}

// handleCreateUser to handle a create user request
func (s *Controller) handleCreateUser(c *gin.Context) {

	var user entity.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	userID, createErr := s.userService.CreateUser(user)
	if createErr != nil {
		c.JSON(createErr.StatusCode, createErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": userID})
}

// handleUpdateUser to handle a update user request
func (s *Controller) handleUpdateUser(c *gin.Context) {
	var user entity.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	userID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	user.ID = userID

	resUser, updateErr := s.userService.UpdateUser(user)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}

	vmResponse := userStructToViewmodelResponse(*resUser)

	c.JSON(http.StatusOK, vmResponse)
}

// handleDeleteUser to handle a delete user request
func (s *Controller) handleDeleteUser(c *gin.Context) {

	userID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	deleteErr := s.userService.DeleteUser(userID)
	if deleteErr != nil {
		c.JSON(deleteErr.StatusCode, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

// handleLogin to handle a user login request
func (s *Controller) handleLogin(c *gin.Context) {
	var credentials = entity.LoginRequest{}

	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	_, loginErr := s.userService.LoginUser(credentials)
	if loginErr != nil {
		c.JSON(loginErr.StatusCode, loginErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Ok"})
}

func (s *Controller) getIDParameter(userParamID string) (int64, *resterrors.RestErr) {
	userID, userErr := strconv.ParseInt(userParamID, 10, 64)
	if userErr != nil {
		return 0, resterrors.NewBadRequestError("Param id should be a number")
	}

	return userID, nil
}

func userStructToViewmodelResponse(user entity.User) (vmUser viewmodel.User) {

	vmUser.ID = user.ID
	vmUser.Name = user.Name
	vmUser.Email = user.Email
	vmUser.DocumentNumber = user.DocumentNumber
	vmUser.CountryCode = user.CountryCode
	vmUser.AreaCode = user.AreaCode
	vmUser.PhoneNumber = user.PhoneNumber
	vmUser.Birthdate = user.Birthdate
	vmUser.Gender = user.Gender
	vmUser.Revenue = user.Revenue
	vmUser.Active = user.Active

	return vmUser
}

func userListStructToViewmodelResponse(users []entity.User) (vmUsers []viewmodel.User) {

	var vmUser viewmodel.User

	for i := 0; i < len(users); i++ {
		vmUser.ID = users[i].ID
		vmUser.Name = users[i].Name
		vmUser.Email = users[i].Email
		vmUser.DocumentNumber = users[i].DocumentNumber
		vmUser.CountryCode = users[i].CountryCode
		vmUser.AreaCode = users[i].AreaCode
		vmUser.PhoneNumber = users[i].PhoneNumber
		vmUser.Birthdate = users[i].Birthdate
		vmUser.Gender = users[i].Gender
		vmUser.Revenue = users[i].Revenue
		vmUser.Active = users[i].Active

		vmUsers = append(vmUsers, vmUser)
	}

	return vmUsers
}
