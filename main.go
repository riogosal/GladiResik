package main

import (
	connection "GladiResik/Connection"
	food "GladiResik/Food"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	DB := options.Client().ApplyURI("mongodb://localhost:27017")
	ClientDB := connection.NewDB(DB)
	FoodData := food.Initialize(ClientDB.Client)
	FoodData.PrintAll()

	router := gin.Default()
	connection.Setup(router, ClientDB, FoodData)
	router.Run(":8080")

	ClientDB.DisconnectDB()
}
