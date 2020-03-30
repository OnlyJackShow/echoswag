package route

import (
	"github.com/labstack/echo/v4"

	"echoswag/controller"
)

func Route(e *echo.Echo)*echo.Echo {
	e.GET("/boss/product/brand/list", controller.QueryBrand)
	return e
}

