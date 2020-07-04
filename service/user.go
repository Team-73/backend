package service

import (
	"strings"

	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/utils/cryptoutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type userService struct {
	svc *Service
}

//newUserService return a new instance of the service
func newUserService(svc *Service) contract.UserService {
	return &userService{
		svc: svc,
	}
}

func (s *userService) GetUsers() (*[]entity.User, *resterrors.RestErr) {

	users, err := s.svc.db.User().GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) GetUserByID(userID int64) (*entity.User, *resterrors.RestErr) {
	user := &entity.User{
		ID: userID,
	}

	user, err := s.svc.db.User().GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUser(user entity.User) (int64, *resterrors.RestErr) {
	if err := user.Validate(); err != nil {
		return 0, err
	}

	user.Password = cryptoutils.GetMd5(user.Password)

	newUser, err := s.svc.db.User().Create(user)
	if err != nil {
		return 0, err
	}

	return newUser, nil
}

func (s *userService) UpdateUser(user entity.User) (*entity.User, *resterrors.RestErr) {

	currentUser, err := s.GetUserByID(user.ID)
	if err != nil {
		return nil, err
	}

	if user.Name != "" {
		currentUser.Name = strings.TrimSpace(user.Name)
	}

	if user.Email != "" {
		currentUser.Email = strings.TrimSpace(user.Email)
	}

	if user.CountryCode != "" {
		currentUser.CountryCode = strings.TrimSpace(user.CountryCode)
	}

	if user.AreaCode != "" {
		currentUser.AreaCode = strings.TrimSpace(user.AreaCode)
	}

	if user.PhoneNumber != "" {
		currentUser.PhoneNumber = strings.TrimSpace(user.PhoneNumber)
	}

	if user.Revenue != 0 {
		currentUser.Revenue = user.Revenue
	}

	updatedUser, err := s.svc.db.User().Update(*currentUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *userService) DeleteUser(userID int64) *resterrors.RestErr {
	return s.svc.db.User().Delete(userID)
}

func (s *userService) LoginUser(request entity.LoginRequest) (*entity.User, *resterrors.RestErr) {
	userLogin := &entity.LoginRequest{
		Email:    request.Email,
		Password: cryptoutils.GetMd5(request.Password),
	}

	user, err := s.svc.db.User().GetByEmailAndPassword(*userLogin)
	if err != nil {
		return nil, err
	}
	return user, nil
}
