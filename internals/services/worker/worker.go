package worker

import (
	"context"
	logger "github.com/labstack/gommon/log"
)

type ServiceWorker struct {
	ReaderSvc ServiceReader
	SenderSvc ServiceSender
	Topic     string
}

func NewServiceWorker(readerSvc ServiceReader, senderSVC ServiceSender, topic string) *ServiceWorker {
	return &ServiceWorker{
		ReaderSvc: readerSvc,
		SenderSvc: senderSVC,
		Topic:     topic,
	}
}

func (w *ServiceWorker) Execute(ctx context.Context) error {
	logger.Infof("Worker started")
	productChunks, err := w.ReaderSvc.ReadFileInDirectory(ctx)
	if err != nil {
		logger.Errorf("Invalid to read file from csv:%s", err)
	}
	logger.Infof("product: %+v", productChunks)
	for idx, chunk := range productChunks {
		logger.Infof("Processing chunk #%d with %d records", idx+1, len(chunk))

		// ส่งข้อมูลในแต่ละ chunk
		for _, product := range chunk {
			err := w.SenderSvc.SendKafkaMSG(ctx, "product_topic", product)
			if err != nil {
				logger.Errorf("Failed to send product to Kafka: %+v, error: %v", product, err)
			}
		}
	}

	logger.Infof("Worker finished successfully")
	return nil
}
