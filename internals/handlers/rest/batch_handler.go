package rest

import (
	"context"
	"github.com/NuttayotSukkum/batch_consumer/internals/handlers"
	error2 "github.com/NuttayotSukkum/batch_consumer/internals/models/dto"
	"github.com/NuttayotSukkum/batch_consumer/internals/models/responses"
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/constants"
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/utils"
	"github.com/NuttayotSukkum/batch_consumer/internals/repositories"
	"github.com/NuttayotSukkum/batch_consumer/internals/services"
	"github.com/labstack/echo/v4"
	logger "github.com/labstack/gommon/log"
	"net/http"
)

type BatchHandler struct {
	preProcessSvc   services.PreProcess
	workerSvc       services.ServiceWorker
	BatchHeaderRepo repositories.BatchHeaderDBRepository
}

func NewBatchHandler(preProcessSvc services.PreProcess, workerSvc services.ServiceWorker, batchHeaderRepo repositories.BatchHeaderDBRepository) handlers.Batch {
	return &BatchHandler{
		preProcessSvc:   preProcessSvc,
		workerSvc:       workerSvc,
		BatchHeaderRepo: batchHeaderRepo,
	}
}

func (h *BatchHandler) Initial(c echo.Context) error {
	var response responses.InitialResponse
	ctx := context.Background()
	batchId, batchDate, err := h.preProcessSvc.PreStart(ctx, constants.DirPath)
	if err != nil {
		logger.Infof("Error pre processing %s", err)
		return c.JSON(http.StatusInternalServerError, error2.ResponseErrorBucketIsEmpty())
	}
	response.BatchHeaderId = batchId
	logger.Warnf("Batch date:%v", batchDate)
	response.BatchDate = batchDate.Format(constants.DATE_TIME_FORMATTER)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			logger.Warnf("Worker stopped: %s", ctx.Err())
			return
		default:
			if err := h.workerSvc.Execute(ctx); err != nil {
				logger.Errorf("Error read file: %s", err)
				h.BatchHeaderRepo.UpdateBatchStatus(ctx, batchId, constants.BatchStatusFailed)
			}
			h.BatchHeaderRepo.UpdateBatchStatus(ctx, batchId, constants.BatchStatusSuccess)
			utils.DeleteDirectory(constants.DirPath)
		}
	}(ctx)
	return c.JSON(http.StatusOK, response)
}
