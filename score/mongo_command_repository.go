package score

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoCommandRepository struct {
	db *mongo.Database
}

func NewMongoCommandRepository(db *mongo.Database) CommandRepository {
	return MongoCommandRepository{db: db}
}

func (receiver MongoCommandRepository) Save(score Score) error {
	collection := receiver.db.Collection("scores")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	scoreInsertResult, err := collection.InsertOne(ctx, score)

	if err != nil {
		return err
	}

	logrus.Infof("%s score inserted with id %s and point %d", scoreInsertResult.InsertedID, score.Id, score.Point)
	return nil
}

func (receiver MongoCommandRepository) Update(score Score) (int64, error) {
	collection := receiver.db.Collection("scores")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateResult, err := collection.UpdateOne(ctx, bson.M{"id": score.Id}, bson.D{{"$set", bson.M{"point": score.Point}}})
	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, nil
}
