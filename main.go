package main

import (
	connection "GladiResik/Connections"
	food "GladiResik/Food"
	"context"

	"github.com/gin-gonic/gin"
)

func main() {
	mainCtx := context.Background()
	mongo := connection.NewDB(mainCtx)
	defer mongo.Disconnect(mainCtx)
	database := mongo.Database("GladiResik")
	router := gin.Default()

	FoodData := food.Initialize(database)
	food.Setup(router, FoodData)

	router.Run(":8080")
}
