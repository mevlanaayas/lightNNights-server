package score

import "go.mongodb.org/mongo-driver/mongo"

type MongoQueryRepository struct {
	db *mongo.Database
}

func NewMongoQueryRepository(db *mongo.Database) QueryRepository {
	return MongoQueryRepository{db: db}
}

func (receiver MongoQueryRepository) Get(id string) error {
	return nil
}

func (receiver MongoQueryRepository) List() error {
	return nil
}
