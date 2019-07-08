package custom

import (
	"strconv"

	"github.com/ClanceyLu/echo-api/conf"
	"github.com/labstack/echo/v4"
)

// Context extends echo.Context
type Context struct {
	echo.Context
}

// QueryArray returns a slice of strings for the key
func (c *Context) QueryArray(key string) []string {
	return c.QueryParams()[key]
}

// QueryDefault returns the query param for the key if it exists, otherwise it returns defalutValue
func (c *Context) QueryDefault(key, defalutValue string) string {
	val := c.QueryParam(key)
	if val == "" {
		val = defalutValue
	}
	return val
}

// PageInfo always returns the page and pageSize of the query
// if page and pageSize not exists, it will return default value
func (c *Context) PageInfo() (int, int) {
	appConf := conf.Conf.Sub("app")
	page, err := strconv.Atoi(c.QueryDefault("page", appConf.GetString("page")))
	pageSize, err := strconv.Atoi(c.QueryDefault("pageSize", "10"))
	if err != nil {
		return appConf.GetInt("page"), appConf.GetInt("pageSize")
	}
	return page, pageSize
}
