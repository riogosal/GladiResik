package food

import (
	entity "GladiResik/Entity"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Foods struct {
	foods []entity.Food
}

func Initialize(client *mongo.Client) *Foods {
	var foods []entity.Food
	collection := client.Database("GladiResik").Collection("Food")

	filter := bson.D{}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var food entity.Food
		err := cursor.Decode(&food)
		if err != nil {
			log.Fatal(err)
		}
		foods = append(foods, food)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return &Foods{
		foods: foods,
	}
}

func (item *Foods) ViewAll(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, item.foods)
}

func (item *Foods) PrintAll() {

}
