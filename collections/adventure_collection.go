package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdventureCollection interface {
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
	DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
	FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult
}