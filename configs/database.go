package configs

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mongoOnce   sync.Once
)

// GetMongoClient inicializa (una sola vez) y devuelve el cliente de MongoDB
func GetMongoClient() (*mongo.Client, error) {
	var err error

	mongoOnce.Do(func() {
		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			log.Fatal("MONGODB_URI no está configurado en las variables de entorno")
		}

		clientOptions := options.Client().ApplyURI(uri)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		mongoClient, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			return
		}

		err = mongoClient.Ping(ctx, nil)
	})

	return mongoClient, err
}

func GetMongoDB(client *mongo.Client) *mongo.Database {
	dbName := os.Getenv("MONGODB_DB_NAME")
	if dbName == "" {
		dbName = "mapsnt" 
	}
	return client.Database(dbName)
}
