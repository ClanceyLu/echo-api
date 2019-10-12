package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ClanceyLu/echo-api/pkg/redis"
	"github.com/labstack/echo/v4"
)

// IUpload 定义上传接口
type IUpload interface {
	Upload(c echo.Context) error
}

type uploadController struct{}

// NewUpload 返回 upload
func NewUpload() IUpload {
	return &uploadController{}
}

func (u uploadController) Upload(c echo.Context) error {
	ip := c.RealIP()
	key := fmt.Sprintf("upload-%s", ip)
	redisClient := redis.Connect()
	times, err := redisClient.Get(key).Result()
	log.Printf("times %s", times)
	if err == redis.Nil {
		if err := redisClient.Set(key, 1, time.Minute).Err(); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else if times != "" {
		t, _ := strconv.Atoi(times)
		if t >= 4 {
			return echo.NewHTTPError(200, "每分钟只能上传四次")
		}
		redisClient.Incr(key)
	}
	return c.String(http.StatusOK, "upload")
}
