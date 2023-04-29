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
	UpdateUser(bool, models.User) (*models.User, *utils.RestErr)
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

func (s *usersService) UpdateUser(isPartial bool, user models.User) (*models.User, *utils.RestErr) {
	current := &models.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
