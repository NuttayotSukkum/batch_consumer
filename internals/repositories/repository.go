package repositories

import (
	"context"
	"github.com/NuttayotSukkum/batch_consumer/internals/models/dao"
)

type (
	BatchHeaderDBRepository interface {
		Insert(ctx context.Context, batchHeader dao.BatchHeader) (*dao.BatchHeader, error)
		UpdateBatchStatus(ctx context.Context, id string, batchStatus string) (*dao.BatchHeader, error)
		//UpdateBatchResult(ctx context.Context, id string, batchStatus string, noRecord int, noSuccess int, noFailed int) (models.BatchHeader, dto)
	}
	KafkaProducer interface {
		Producer(ctx context.Context, topic string, message []byte) error
	}
)
