package mongodb

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DefaultMongoDBUsername = ""
	DefaultMongoDBPassword = ""
	DefaultMongoDBHost     = "localhost"
	DefaultMongoDBPort     = "27017"
	DefaultMongoDBDatabase = ""
)

var (
	Client *mongo.Client
)

type Config struct {
	MongoDBUsername string
	MongoDBPassword string
	MongoDBHost     string
	MongoDBPort     string
	MongoDBDatabase string
}

func assignDefaultValues(config *Config) {
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
}

func init() {
	config := Config{
		MongoDBUsername: os.Getenv("MONGODB_USERNAME"),
		MongoDBPassword: os.Getenv("MONGODB_PASSWORD"),
		MongoDBHost:     os.Getenv("MONGODB_HOST"),
		MongoDBPort:     os.Getenv("MONGODB_PORT"),
		MongoDBDatabase: os.Getenv("MONGODB_DATABASE"),
	}

	assignDefaultValues(&config)

	u := &url.URL{
		Scheme: "mongodb",
		Host:   fmt.Sprintf("%s:%s", config.MongoDBHost, config.MongoDBPort),
		Path:   fmt.Sprintf("/%s", config.MongoDBDatabase),
	}
	if config.MongoDBUsername != "" && config.MongoDBPassword != "" {
		u.User = url.UserPassword(config.MongoDBUsername, config.MongoDBPassword)
	}
	connectionString := u.String()

	clientOptions := options.Client().ApplyURI(connectionString)

	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = Client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
}
