package food

import (
	entity "GladiResik/Entity"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Foods struct {
	foods mongo.Collection
}

func Initialize(client *mongo.Database) *Foods {
	return &Foods{
		foods: *client.Collection("foods"),
	}
}

func (repo *Foods) ViewAll(ctx context.Context) ([]entity.Food, error) {
	var foods []entity.Food

	filter := bson.D{}
	cursor, err := repo.foods.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var food entity.Food
		err := cursor.Decode(&food)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	// Tidak perlu di cek sih ini
	// if err := cursor.Err(); err != nil {
	// 	return nil, err
	// }

	return foods, nil
}
