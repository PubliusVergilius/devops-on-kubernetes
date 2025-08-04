package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lock = &sync.Mutex{}
type singleConnection struct {
	client *mongo.Client
}

var instance *singleConnection

func getConnectionInstance(ctx context.Context, uri string) *singleConnection {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
			if err != nil {
				log.Fatal(err)
			}

			err = client.Ping(ctx, nil)
			if err != nil {
				log.Fatal("Could not stabilish a successfull connection with Mongo database")
			}
			log.Panicf("MongoDB connection successfull!")
			instance = &singleConnection{
				client,
			}
		} else {
			fmt.Println("MongoDB server already connected")
		}
	} else {
		fmt.Println("MongoDB server already connected")
	}
	return instance
}

const defaultUri string = "mongodb://localhost:27017"

func Connect(dbName string) (*mongo.Database, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Printf("Using default MONGODB_URI value: %s\n", defaultUri)
		uri = defaultUri
	}

	conn := getConnectionInstance(ctx, uri)
	
	// defer client.Disconnect(ctx)

	return conn.client.Database(dbName), ctx
}
