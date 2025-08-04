package collections

import (
	"api/config"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func ColorCollection() (*mongo.Collection, context.Context) {
	db, ctx := config.Connect("colorsDb")
	return db.Collection("colors"), ctx
}
