package handlers

import "github.com/labstack/echo/v4"

type (
	Batch interface {
		Initial(c echo.Context) error
	}
)
