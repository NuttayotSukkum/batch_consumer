package rest

import (
	"context"
	"github.com/NuttayotSukkum/batch_consumer/internals/handlers"
	error2 "github.com/NuttayotSukkum/batch_consumer/internals/models/error"
	"github.com/NuttayotSukkum/batch_consumer/internals/models/responses"
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/constants"
	"github.com/NuttayotSukkum/batch_consumer/internals/services"
	"github.com/labstack/echo/v4"
	logger "github.com/labstack/gommon/log"
	"net/http"
)

type BatchHandler struct {
	preProcessSvc services.PreProcess
}

func NewBatchHandler(preProcessSvc services.PreProcess) handlers.Batch {
	return &BatchHandler{
		preProcessSvc: preProcessSvc,
	}
}

func (h *BatchHandler) Initial(c echo.Context) error {
	var response responses.InitialResponse
	ctx := context.Background()
	batchId, batchDate, err := h.preProcessSvc.PreStart(ctx, constants.DirPath)
	if err != nil {
		logger.Infof("Error pre processing %s", err)
		return c.JSON(http.StatusInternalServerError, error2.GenericError)
	}
	response.BatchHeaderId = batchId
	logger.Warnf("Batch date:%v", batchDate)
	response.BatchDate = batchDate.Format(constants.DATE_TIME_FORMATTER)
	return c.JSON(http.StatusOK, response)
}
