package service

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"os"
	"scanner/internal/entity"
	"scanner/pkg/connection"
	"scanner/pkg/env"
)

var (
	gormClient  *gorm.DB
	mongoClient *mongo.Client
	redisClient *redis.Client

	authService *AuthService
)

func init() {
	if err := env.LoadEnv(".test.env"); err != nil {
		panic(err)
	}

	mongoClient = connection.NewMongoConnection(os.Getenv("MONGO_URI"))

	gormClient = connection.NewGormConnection(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("MYSQL_HOST"),
	)

	gormClient.AutoMigrate(&entity.User{})

	redisClient = connection.NewRedisConnection(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("REDIS_DB"),
	)

	authService = NewAuthService(gormClient, redisClient)
}
