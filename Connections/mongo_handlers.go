package connection

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	ClientOpt *options.ClientOptions
	Client    *mongo.Client
}

func NewDB(clientOptions *options.ClientOptions) *MongoDB {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("connect")
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Ping")
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &MongoDB{
		ClientOpt: clientOptions,
		Client:    client,
	}
}

func (MDB *MongoDB) DisconnectDB() {
	err := MDB.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection closed.")
}
