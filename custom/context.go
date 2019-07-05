package custom

import "github.com/labstack/echo/v4"

// Context extends echo.Context
type Context struct {
	echo.Context
}

// QueryArray returns an array
func (c *Context) QueryArray(key string) []string {
	return c.QueryParams()[key]
}
