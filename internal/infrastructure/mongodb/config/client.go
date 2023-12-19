package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfigInterface interface {
	GetClient(ctx context.Context) (*mongo.Client, error)
}

type MongoDBConfig struct {
	uri string
}

func NewMongoDBConfig(uri string) MongoDBConfigInterface {
	return &MongoDBConfig{
		uri: uri,
	}
}

func (c *MongoDBConfig) getConnectionString() string {
	return c.uri
}

func (c *MongoDBConfig) GetClient(ctx context.Context) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(c.getConnectionString()).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}
