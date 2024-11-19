package preprocess

import (
	"context"
	"github.com/NuttayotSukkum/batch_consumer/internals/models/dao"
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/constants"
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/utils"
	"github.com/NuttayotSukkum/batch_consumer/internals/repositories"
	"github.com/NuttayotSukkum/batch_consumer/internals/services/clients"
	logger "github.com/labstack/gommon/log"
	"time"
)

type PreProcess struct {
	BatchHeaderDBRepo repositories.BatchHeaderDBRepository
	S3Client          clients.S3Client
}

func NewPreProcessService(batchHeaderDBRepo repositories.BatchHeaderDBRepository, S3Client clients.S3Client) PreProcess {
	return PreProcess{
		BatchHeaderDBRepo: batchHeaderDBRepo,
		S3Client:          S3Client,
	}
}

func (svc *PreProcess) PreStart(ctx context.Context, dir string) (string, time.Time, error) {
	startDate := time.Now()
	logger.Infof("%v Start running job . . . start date %v  ", ctx, startDate)
	utils.EmptyInputDirectory()
	batchHeader := (&dao.BatchHeader{}).BuildBatchHeader(startDate.Format(constants.DATE_FORMATTER)+constants.FILE_NAME, constants.BatchStatusProcessing)
	batch, err := svc.BatchHeaderDBRepo.Insert(ctx, *batchHeader)
	if err != nil {
		logger.Errorf("ctx: %v Error to create batch header: %s ", ctx, err)
		return "", batch.BatchDate, err
	}
	if err := svc.S3Client.DownloadFile(ctx, dir); err != nil {
		logger.Errorf("Error downloading files from S3: %s", err)
		svc.BatchHeaderDBRepo.UpdateBatchStatus(ctx, batch.Id, constants.BatchStatusFailed)
		return batch.Id, batchHeader.BatchDate, err
	}
	logger.Infof("ctx: %v All files downloaded successfully to %s", ctx, dir)
	return batch.Id, batchHeader.BatchDate, err
}
