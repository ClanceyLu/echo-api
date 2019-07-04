package custom

import "github.com/labstack/echo/v4"

// Binder custom binder
type Binder struct{}

// Bind implementation echo bind interface
func (binder *Binder) Bind(i interface{}, c echo.Context) error {
	db := new(echo.DefaultBinder)
	if err := db.Bind(i, c); err != nil && err != echo.ErrUnsupportedMediaType {
		return err
	}
	// 验证
	if err := c.Validate(i); err != nil {
		return echo.NewHTTPError(422, err.Error()).SetInternal(err)
	}
	return nil
}
