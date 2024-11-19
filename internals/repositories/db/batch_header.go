package db

import (
	"context"
	"github.com/NuttayotSukkum/batch_consumer/internals/models/dao"
	"github.com/NuttayotSukkum/batch_consumer/internals/repositories"
	logger "github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type batchHeader struct {
	db *gorm.DB
}

func NewBatchHeaderRepository(db *gorm.DB) repositories.BatchHeaderDBRepository {
	return &batchHeader{db: db}
}

func (repo *batchHeader) Insert(ctx context.Context, batchHeader dao.BatchHeader) (*dao.BatchHeader, error) {
	if err := repo.db.Create(&batchHeader).Error; err != nil {
		logger.Errorf("%v: unable to create batch header: %s", ctx, err)
		return nil, err
	}
	return &batchHeader, nil
}

func (repo *batchHeader) UpdateBatchStatus(ctx context.Context, id, batchStatus string) (*dao.BatchHeader, error) {
	if err := repo.db.Model(&dao.BatchHeader{}).
		Where(&dao.BatchHeader{Id: id}).
		Updates(&dao.BatchHeader{
			BatchStatus: batchStatus,
		}).Error; err != nil {
		logger.Errorf("%v :unable to update batch header status : %s", ctx, err)
	}
	return &dao.BatchHeader{}, nil
}
