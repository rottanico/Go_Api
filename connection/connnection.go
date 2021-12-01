package connection

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB       = "arcoiris"
	Product  = "productos"
	Category = "categoria"
)

func GetEnv(key, defaultValue string) string {

	value, defined := os.LookupEnv(key)
	if !defined {
		return defaultValue
	}

	return value
}

func GetCollection(collectionName string) *mongo.Collection {
	connection := GetEnv("DB_CONN", "mongodb://localhost:27017")
	client, err := mongo.NewClient(options.Client().ApplyURI(connection))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return client.Database(DB).Collection(collectionName)
}
