package main

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"os"
	"scanner/cmd/api/internal/app"
	"scanner/cmd/api/internal/app/handler"
	"scanner/cmd/api/internal/container"
	"scanner/pkg/connection"
	"scanner/pkg/env"
)

var (
	mongoClient *mongo.Client
	gormClient  *gorm.DB
	redisClient *redis.Client
)

func init() {
	if err := env.LoadEnv(".env"); err != nil {
		panic(err)
	}

	mongoClient = connection.NewMongoConnection(os.Getenv("MONGO_URI"))
	gormClient = connection.NewGormConnection(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("MYSQL_HOST"),
	)
	redisClient = connection.NewRedisConnection(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("REDIS_DB"),
	)
}

func main() {
	_container := container.NewContainer(gormClient, redisClient, mongoClient)
	_app := app.NewApp(_container)
	_app.AppendHandler(handler.Handlers(_container)...)
	_app.SetRouter()
	if err := _app.Run(); err != nil {
		panic(err)
	}
}
