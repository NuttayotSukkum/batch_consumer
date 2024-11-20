package dao

import (
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/utils"
	"github.com/google/uuid"
	"time"
)

type BatchHeader struct {
	Id          string `gorm:`
	BatchName   string
	BatchDate   time.Time
	BatchStatus string
	Reason      *string
	NoRecords   *int
	NoSuccess   *int
	NoFailed    *int
}

func (BatchHeader) TableName() string {
	return "tbl_batch_headers"
}

func (b *BatchHeader) BuildBatchHeader(batchName, batchStatus string) *BatchHeader {
	return &BatchHeader{
		Id:          uuid.New().String(),
		BatchName:   batchName,
		BatchDate:   utils.TimeLocal(time.Now()),
		BatchStatus: batchStatus,
		Reason:      nil,
		NoRecords:   nil,
		NoSuccess:   nil,
		NoFailed:    nil,
	}
}
