package mongodb

import (
	"context"
	"fmt"

	getEnv "social-network/infra/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBProvider struct {
	client *mongo.Client
}

func Connect() (*MongoDBProvider, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(getEnv.Execute("MONGODB_URL", "mongodb://localhost:27017/social-network")))
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %w", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	return &MongoDBProvider{
		client: client,
	}, nil
}

func (p *MongoDBProvider) Collection(name string) *mongo.Collection {
	return p.client.Database("social-network").Collection(name)
}

func (p *MongoDBProvider) Close() error {
	return p.client.Disconnect(context.Background())
}
