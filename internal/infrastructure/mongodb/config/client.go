package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	Host string
	Port string
	User string
	Pass string
}

func NewMongoDBConfig() *MongoDBConfig {
	return &MongoDBConfig{}
}

func (c *MongoDBConfig) GetHost() string {
	return c.Host
}

func (c *MongoDBConfig) GetPort() string {
	return c.Port
}

func (c *MongoDBConfig) GetUser() string {
	return c.User
}

func (c *MongoDBConfig) GetPass() string {
	return c.Pass
}

func (c *MongoDBConfig) GetConnectionString() string {
	return "mongodb://" + c.GetUser() + ":" + c.GetPass() + "@" + c.GetHost()
}

func (c *MongoDBConfig) GetConnection(ctx context.Context) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(c.GetConnectionString()).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}
