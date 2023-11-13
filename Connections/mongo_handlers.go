package connection

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type MongoDB struct {
// 	ClientOpt *options.ClientOptions
// 	Client
// }

func NewDB(mainCtx context.Context) *mongo.Client {
	client, err := mongo.Connect(mainCtx, options.Client().ApplyURI(fmt.Sprintf(
		"mongodb://felix:123456@localhost:27017/?directConnection=true&authSource=admin",
	)))
	if err != nil {
		fmt.Println("connect")
		log.Fatal(err)
	}

	err = client.Ping(mainCtx, nil)
	if err != nil {
		fmt.Println("Ping")
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// func (MDB *MongoDB) DisconnectDB() {
// 	err := MDB.Client.Disconnect(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connection closed.")
// }
