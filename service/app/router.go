package app

import (
	"github.com/ClanceyLu/echo-api/service"
	"github.com/go-redis/redis/v7"

	"github.com/ClanceyLu/echo-api/service/app/upload"
	"github.com/ClanceyLu/echo-api/service/app/user"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type app struct {
	db          *gorm.DB
	redisClient *redis.Client
}

// New 新建一个 app 服务
func New(db *gorm.DB, redisClient *redis.Client) service.Service {
	return &app{db, redisClient}
}

// Router 注册 app 路由
func (app *app) Router(r *echo.Group) {
	appRouter := r.Group("/app")

	user := user.New(app.db)
	user.Router(appRouter)

	upload := upload.New(app.redisClient)
	upload.Router(appRouter)
}