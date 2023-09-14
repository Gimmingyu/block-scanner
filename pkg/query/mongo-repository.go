package query

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository[T Collection] struct {
	conn *mongo.Collection
}

func (m *mongoRepository[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var (
		result = new(T)
		err    error
	)

	if err = m.conn.FindOne(ctx, bson.D{{"_id", id}}).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mongoRepository[T]) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*T, error) {
	var (
		result = new(T)
		err    error
	)

	if err = m.conn.FindOne(ctx, filter, opts...).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mongoRepository[T]) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*T, error) {
	var (
		result = make([]*T, 0)
		cursor *mongo.Cursor
		err    error
	)

	if cursor, err = m.conn.Find(ctx, filter, opts...); err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mongoRepository[T]) Insert(ctx context.Context, document *T, opts ...*options.InsertOneOptions) error {
	var err error

	if _, err = m.conn.InsertOne(ctx, document, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) InsertMany(ctx context.Context, documents []*T, opts ...*options.InsertManyOptions) error {
	var (
		slice []interface{}
		err   error
	)

	for _, document := range documents {
		slice = append(slice, document)
	}

	if _, err = m.conn.InsertMany(ctx, slice, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	var (
		err error
	)

	if _, err = m.conn.UpdateOne(ctx, filter, update, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	var (
		err error
	)

	if _, err = m.conn.UpdateMany(ctx, filter, update, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) UpdateByID(ctx context.Context, id uint, update interface{}, opts ...*options.UpdateOptions) error {
	var (
		err error
	)

	if _, err = m.conn.UpdateByID(ctx, id, update, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error {
	var (
		err error
	)

	if _, err = m.conn.DeleteOne(ctx, filter, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error {
	var (
		err error
	)

	if _, err = m.conn.DeleteMany(ctx, filter, opts...); err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository[T]) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) ([]*T, error) {
	var (
		result []*T
		cursor *mongo.Cursor
		err    error
	)

	if cursor, err = m.conn.Aggregate(ctx, pipeline, opts...); err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *mongoRepository[T]) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	var (
		result int64
		err    error
	)

	if result, err = m.conn.CountDocuments(ctx, filter, opts...); err != nil {
		return 0, err
	}

	return result, nil
}
