package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoDbClient(cfg *Config) (*mongo.Database, *mongo.Client) {
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port)
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(cfg.Database), client
}

func DisposeClient(client *mongo.Client) {
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}
