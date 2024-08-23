package mongodb

import (
	"booking_service/config"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(cfg *config.Config) (*mongo.Client, *mongo.Database, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongo:27017").SetAuth(options.Credential{
		Username: "booking_service",
		Password: "pass",
	}))
	
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}

	db := client.Database(cfg.MongoDBName)
	return client, db, nil
}
