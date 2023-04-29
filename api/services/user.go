package services

import (
	"amaranth/api/models"
	"amaranth/api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	CreateUser(models.User) (*models.User, *utils.RestErr)
	GetUser(primitive.ObjectID) (*models.User, *utils.RestErr)
}

type usersService struct{}

func (s *usersService) CreateUser(user models.User) (*models.User, *utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) GetUser(userId primitive.ObjectID) (*models.User, *utils.RestErr) {
	result := &models.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
