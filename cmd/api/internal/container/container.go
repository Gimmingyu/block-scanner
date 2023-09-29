package container

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Container struct {
	db    *gorm.DB
	redis *redis.Client
	mongo *mongo.Client
}

func NewContainer(gormClient *gorm.DB, redisClient *redis.Client, mongoClient *mongo.Client) *Container {
	return &Container{
		db:    gormClient,
		redis: redisClient,
		mongo: mongoClient,
	}
}
