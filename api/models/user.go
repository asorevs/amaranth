package models

import (
	"amaranth/api/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	FirstName   string             `json:"name"`
	LastName    string             `json:"lastName"`
	Email       string             `json:"email"`
	DateCreated string             `json:"creationDate"`
	Status      string             `json:"status"`
}

type Users []User

func (user *User) Validate() *utils.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return utils.NewBadRequestError("invalid email address")
	}

	return nil
}
