package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection interface {
}

type MongoRepository[T Collection] interface {
	FindByID(ctx context.Context, id uint) (*T, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*T, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*T, error)
	Insert(ctx context.Context, document *T, opts ...*options.InsertOneOptions) error
	InsertMany(ctx context.Context, documents []*T, opts ...*options.InsertManyOptions) error
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error
	UpdateByID(ctx context.Context, id uint, update interface{}, opts ...*options.UpdateOptions) error
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error
	Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) ([]*T, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
}

func NewMongoRepository[T Collection](conn *mongo.Collection) MongoRepository[T] {
	return &mongoRepository[T]{conn}
}

type Table interface {
	Table() string
}
