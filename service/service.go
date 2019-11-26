package service

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service 定义了服务接口
type Service interface {
	Router(r *echo.Group)
}

// Controller 定义了 controller 的结构
type Controller struct {
	Mysql *gorm.DB
	Redis *redis.Client
	Mongo *mongo.Client
}
