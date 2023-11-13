package main

import (
	connection "GladiResik/Connections"
	food "GladiResik/Food"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	mainCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Mongo Initiliazation
	ClientDB := connection.NewDB(mainCtx)
	defer ClientDB.Disconnect(mainCtx)
	router := gin.Default()
	db := ClientDB.Database("GladiResik")

	FoodData := food.Initialize(db)
	food.Setup(router, FoodData)

	// FoodData.PrintAll()

	// connection.Setup(router, ClientDB, FoodData)
	router.Run(":12345")
}
