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

	return foods, nil
}

func (repo *Foods) InsertFood(ctx context.Context, item entity.Food) {
	repo.foods.InsertOne(ctx, item)
}

func (repo *Foods) GetFood(ctx context.Context, id int) entity.Food {
	cursor := repo.foods.FindOne(ctx, bson.M{"id": id})
	var result entity.Food
	err := cursor.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (repo *Foods) Update(ctx context.Context, id int) {
	repo.foods.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": bson.M{"availabe": true}},
	)
}
