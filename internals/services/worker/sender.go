package worker

import (
	"context"
	"encoding/json"
	"github.com/NuttayotSukkum/batch_consumer/internals/models/dao"
	"github.com/NuttayotSukkum/batch_consumer/internals/repositories"
	logger "github.com/labstack/gommon/log"
)

type ServiceSender struct {
	kafkaProducer repositories.KafkaProducer
}

func NewServiceSender(kafkaProducer repositories.KafkaProducer) *ServiceSender {
	return &ServiceSender{
		kafkaProducer: kafkaProducer,
	}
}

func (s *ServiceSender) SendKafkaMSG(ctx context.Context, topic string, product dao.Product) error {
	msg, err := json.Marshal(product)
	if err != nil {
		logger.Panicf("Failed to marshal product: %+v, error: %v", product, err)
		return err
	}
	if err := s.kafkaProducer.Producer(ctx, topic, msg); err != nil {
		logger.Errorf("Failed to send message to kafka, topic: %s, error : %s", topic, err)
		logger.Panicf("Service is down . . .")
		return err
	}

	return nil
}
