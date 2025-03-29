package connections

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"medtracker/medtracker/internal/app/config"
)

type Connections struct {
	DB *mongo.Client
}

func NewConnections(cfg *config.Config) (*Connections, error) {
	clientOptions := options.Client().ApplyURI(cfg.DB.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Проверим подключение
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	return &Connections{DB: client}, nil
}

func (c *Connections) Close() {
	if err := c.DB.Disconnect(context.Background()); err != nil {
		log.Printf("Error closing MongoDB connection: %v", err)
	}
}
