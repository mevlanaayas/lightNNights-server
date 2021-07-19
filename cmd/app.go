package cmd

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"lightNNights/api"
	"lightNNights/config"
	"lightNNights/score"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Run() error {
	c := &config.Config{
		Port: 9000,
	}
	err := c.InitMongoUrl()
	if err != nil {
		return err
	}

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = mongoClient.Disconnect(ctx)
		if err != nil {
			logrus.Error(err)
		}
	}()

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	db := mongoClient.Database("user")

	mongoCommandRepository := score.NewMongoCommandRepository(db)
	mongoQueryRepository := score.NewMongoQueryRepository(db)
	commandService := score.NewCommandService(mongoCommandRepository)
	queryService := score.NewQueryService(mongoQueryRepository)

	scoreHandler := score.NewHandler(commandService, queryService)

	server := api.NewAPI(c, scoreHandler)

	if err := server.Start(); err != nil {
		return err
	}
	return nil
}
