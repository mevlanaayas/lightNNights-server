package score

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoQueryRepository struct {
	db *mongo.Database
}

func NewMongoQueryRepository(db *mongo.Database) QueryRepository {
	return MongoQueryRepository{db: db}
}

func (receiver MongoQueryRepository) Get(id string) error {
	return nil
}

func (receiver MongoQueryRepository) List() ([]Score, error) {

	collection := receiver.db.Collection("scores")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var scores []Score
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"point", -1}})
	findOptions.SetLimit(50)
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &scores)
	if err != nil {
		return nil, err
	}
	return scores, nil
}
