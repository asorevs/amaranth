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

	db := mongodb.Client.Database(mongodb.Config.MongoDBDatabase)
	collection := db.Collection(mongodb.Config.MongoDBCollection)

	result := collection.FindOne(context.Background(), bson.M{"_id": user.Id})
	if result.Err() == mongo.ErrNoDocuments {
		return utils.NewNotFoundError(fmt.Sprintf("user %s not found", user.Id.Hex()))
	} else if result.Err() != nil {
		return utils.NewDBError("error when getting user from DB")
	}

	if err := result.Decode(&user); err != nil {
		return utils.NewDBError("error decoding user from DB")
	}

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
		return utils.NewDBError("error when validating duplicate emails")
	}

	user.Id = primitive.NewObjectID()
	user.DateCreated = utils.GetNowString()

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return utils.NewDBError("error when saving user to database")
	}
	return nil
}
