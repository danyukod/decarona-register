package main

import (
	"context"
	"github.com/danyukod/decarona-register/configs"
	"github.com/danyukod/decarona-register/configs/logger"
	"github.com/danyukod/decarona-register/internal/infrastructure/mongodb/config"
)

func main() {
	ctx := context.Background()
	defer ctx.Done()
	logger.Info("About to start User API...")
	conf, err := configs.LoadConfig("cmd/server")
	if err != nil {
		logger.Error("Failed to load config: ", err)
		return
	}
	mongoDbConf := config.NewMongoDBConfig(conf.GetMongoDbUri())

	client, err := mongoDbConf.GetClient(ctx)
	if err != nil {
		logger.Error("Failed to connect to MongoDB: ", err)
		return
	}

	err = client.Disconnect(ctx)
	if err != nil {
		return
	}

}
