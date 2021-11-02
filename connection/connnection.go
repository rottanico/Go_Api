package connection

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "arcoiris"
	Product          = "productos"
	Category         = "categoria"
)

func GetCollection(collectionName string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return client.Database(DB).Collection(collectionName)
}
