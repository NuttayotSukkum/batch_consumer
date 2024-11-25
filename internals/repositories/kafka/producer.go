package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/NuttayotSukkum/batch_consumer/internals/repositories"
	logger "github.com/labstack/gommon/log"
)

type ServiceProducer struct {
	producer sarama.SyncProducer
}

func NewServiceProducer(p sarama.SyncProducer) repositories.KafkaProducer {
	return &ServiceProducer{producer: p}
}

func (producer *ServiceProducer) Producer(ctx context.Context, topic string, message []byte) error {
	logger.Infof("....Start send message kafka")
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	logger.Infof("ctx: %v  Attempting to send Kafka message: %+v", topic)
	partition, offset, err := producer.producer.SendMessage(kafkaMsg)
	if err != nil {
		logger.Errorf("failed to send Kafka message:%s", err)
	}

	logger.Infof("ctx: %v Kafka message sent with trace: topic : %s partition: %voffset: %v", ctx, topic, partition, offset)
	return nil
}
