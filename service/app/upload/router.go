package upload

import (
	"github.com/ClanceyLu/echo-api/service"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
)

type uploadService struct {
	redis *redis.Client
}

// New 返回一个 upload 服务
func New(redisClient *redis.Client) service.Service {
	return &uploadService{redisClient}
}

func (u *uploadService) Router(r *echo.Group) {
	r.POST("/upload", u.upload)
}
