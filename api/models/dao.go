package models

import (
	"amaranth/api/datasources/mongodb"
	"amaranth/api/utils"
	"context"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *utils.RestErr {
	err := mongodb.Client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return utils.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	user.Status = result.Status

	return nil
}

func (user *User) Save() *utils.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return utils.NewBadRequestError(fmt.Sprintf("email %s is already registered", user.Email))
		}
		return utils.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
