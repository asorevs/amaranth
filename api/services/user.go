package services

import "amaranth/api/models"

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	CreateUser(models.User) (*models.User, error)
	GetUser(int64) (*models.User, error)
}

type usersService struct{}

func (s *usersService) CreateUser(user models.User) (*models.User, error) {
	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*models.User, error) {
	return &models.User{Id: userId}, nil
}
