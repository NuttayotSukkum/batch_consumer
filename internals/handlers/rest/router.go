package rest

import (
	"context"
	"github.com/NuttayotSukkum/batch_consumer/internals/repositories"
	"github.com/NuttayotSukkum/batch_consumer/internals/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func InitRouter(ctx context.Context, preProcess services.PreProcess, worker services.ServiceWorker, bathHeaderRepo repositories.BatchHeaderDBRepository) *echo.Echo {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, echo.Map{"message": "Service is Running !!"}, " ")
	})

	api := e.Group("/api")
	batch := api.Group("/batch")
	v1 := batch.Group("/v1")
	v1.Use(middleware.Recover())
	handler := NewBatchHandler(preProcess, worker, bathHeaderRepo)
	v1.POST("/initialize", handler.Initial)
	return e
}
