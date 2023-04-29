package models

import (
	"amaranth/api/datasources/mongodb"
	"amaranth/api/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	usersDB = make(map[primitive.ObjectID]*User)
)

func (user *User) Get() *utils.RestErr {
	err := mongodb.Client.Ping(context.Background(), nil)
	if err != nil {
		utils.NewDBError("encountered an error while attempting to ping the database server")
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
	db := mongodb.Client.Database(mongodb.Config.MongoDBDatabase)
	collection := db.Collection(mongodb.Config.MongoDBCollection)

	existingUser := User{}
	err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return utils.NewBadRequestError(fmt.Sprintf("email %s is already registered", user.Email))
	} else if err != mongo.ErrNoDocuments {
		return utils.NewDBError("error when validating duplicated emails")
	}

	user.Id = primitive.NewObjectID()
	user.DateCreated = utils.GetNowString()

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return utils.NewDBError("error when saving user to database")
	}
	usersDB[user.Id] = user
	return nil
}
