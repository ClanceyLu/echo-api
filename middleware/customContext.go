package middleware

import (
	"github.com/ClanceyLu/echo-api/custom"
	"github.com/labstack/echo/v4"
)

// CustomContext returns a CustomContext middleware
// CustomContext middleware have many custom functions
func CustomContext() func(h echo.HandlerFunc) echo.HandlerFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &custom.Context{
				Context: c,
			}
			return h(cc)
		}
	}
}
