package app

import (
	"github.com/ClanceyLu/echo-api/service"
	"github.com/ClanceyLu/echo-api/service/app/post"
	"github.com/ClanceyLu/echo-api/service/app/upload"
	"github.com/ClanceyLu/echo-api/service/app/user"

	"github.com/labstack/echo/v4"
)

type app service.Controller

// New 新建一个 app 服务
func New(server *service.Controller) service.Service {
	var app app = app(*server)
	return &app
}

// Router 注册 app 路由
func (app *app) Router(r *echo.Group) {
	appRouter := r.Group("/app")

	{
		user := user.New(service.Controller(*app))
		/**
		 * @api {get} /app/user 用户列表
		 * @apiName GetUsers
		 * @apiGroup User
		 * @apiParam {Number} page 页数
		 * @apiParam {Number} pageSize 分页
		 * @apiSuccessExample {json} body:
		 * {
		 *	"list": [],
		 *	"count": 4
		 * }
		 */
		appRouter.GET("/user", user.GetUsers)
		/**
		 * @api {get} /app/user/:id 用户详情
		 * @apiName GetUser
		 * @apiParam {Number} id 用户 ID
		 * @apiSuccessExample {json} body:
		 * {
		 *	"id": 1,
		 *	"nickname": "xiaoming"
		 * }
		 */
		appRouter.GET("/user/:id", user.GetUser)
	}

	{
		post := post.New(service.Controller(*app))
		appRouter.POST("/post", post.PostPost)
		appRouter.GET("/post", post.GetPosts)
	}

	upload := upload.New(app.Redis)
	upload.Router(appRouter)
}
