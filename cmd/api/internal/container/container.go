package container

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"scanner/cmd/api/internal/app/service"
)

type Container struct {
	authService *service.AuthService
}

func NewContainer(gormClient *gorm.DB, redisClient *redis.Client, mongoClient *mongo.Client) *Container {
	return &Container{
		authService: service.NewAuthService(gormClient, redisClient),
	}
}

func (c *Container) AuthService() *service.AuthService {
	return c.authService
}
