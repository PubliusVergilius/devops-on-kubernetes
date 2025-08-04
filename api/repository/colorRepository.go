package repository

import (
	"api/collections"
	"api/schemas"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Save(colorSchema schemas.ColorSchema) string {
	col, ctx := collections.ColorCollection()

	filter := bson.M{colorSchema.Key: colorSchema.Value}

	var color *schemas.ColorSchema
	err := col.FindOne(ctx, filter).Decode(&color)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No color with that name")
			return ""
		}
		log.Fatal(err)
	}

	if color != nil {
		return color.Value
	}

	return os.Getenv("DEFAULT_COLOR")

}
