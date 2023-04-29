package mongodb

import (
	"amaranth/api/utils"
	"context"
	"fmt"
	"net/url"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DefaultMongoDBUsername   = ""
	DefaultMongoDBPassword   = ""
	DefaultMongoDBHost       = "localhost"
	DefaultMongoDBPort       = "27017"
	DefaultMongoDBDatabase   = "amaranth"
	DefaultMongoDBCollection = "users"
)

var (
	Client *mongo.Client
	Config *config
)

type config struct {
	MongoDBUsername   string
	MongoDBPassword   string
	MongoDBHost       string
	MongoDBPort       string
	MongoDBDatabase   string
	MongoDBCollection string
}

func assignDefaultValues(config *config) {
	if config.MongoDBUsername == "" {
		config.MongoDBUsername = DefaultMongoDBUsername
	}
	if config.MongoDBPassword == "" {
		config.MongoDBPassword = DefaultMongoDBPassword
	}
	if config.MongoDBHost == "" {
		config.MongoDBHost = DefaultMongoDBHost
	}
	if config.MongoDBPort == "" {
		config.MongoDBPort = DefaultMongoDBPort
	}
	if config.MongoDBDatabase == "" {
		config.MongoDBDatabase = DefaultMongoDBDatabase
	}
	if config.MongoDBCollection == "" {
		config.MongoDBCollection = DefaultMongoDBCollection
	}
}

func init() {
	Config = &config{
		MongoDBUsername:   os.Getenv("MONGODB_USERNAME"),
		MongoDBPassword:   os.Getenv("MONGODB_PASSWORD"),
		MongoDBHost:       os.Getenv("MONGODB_HOST"),
		MongoDBPort:       os.Getenv("MONGODB_PORT"),
		MongoDBDatabase:   os.Getenv("MONGODB_DATABASE"),
		MongoDBCollection: os.Getenv("MONGODB_COLLECTION"),
	}

	assignDefaultValues(Config)

	u := &url.URL{
		Scheme: "mongodb",
		Host:   fmt.Sprintf("%s:%s", Config.MongoDBHost, Config.MongoDBPort),
		Path:   fmt.Sprintf("/%s", Config.MongoDBDatabase),
	}
	if Config.MongoDBUsername != "" && Config.MongoDBPassword != "" {
		u.User = url.UserPassword(Config.MongoDBUsername, Config.MongoDBPassword)
	}
	connectionString := u.String()

	clientOptions := options.Client().ApplyURI(connectionString)

	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		utils.NewDBError("failed to establish a connection with the database")
	}

	err = Client.Ping(context.Background(), nil)
	if err != nil {
		utils.NewDBError("encountered an error while attempting to ping the database server")
	}
	fmt.Println("Connected to MongoDB!")
}
