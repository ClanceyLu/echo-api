package service

import (
	"github.com/ClanceyLu/echo-api/pkg/mongo"
	"github.com/ClanceyLu/echo-api/pkg/mysql"
	"github.com/ClanceyLu/echo-api/pkg/redis"
	"github.com/labstack/echo/v4"
)

// Mysql client
var Mysql mysql.Client

// Redis client
var Redis redis.Client

// Mongo client
var Mongo mongo.Client

func init() {
	Mysql = mysql.Connect()
	Redis = redis.Connect()
	Mongo = mongo.Connect()
}

// Service 定义了服务接口
type Service interface {
	Router(r *echo.Group)
}
