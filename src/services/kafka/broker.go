package kafka

import (
	brokersettings "app/src/app_settings/brokers"
	"app/src/core/brokers/common"
)

type KafkaBroker struct {
	producer *Producer
	consumer *Consumer
}

func NewKafkaBroker(settings *brokersettings.KafkaSettings) *KafkaBroker {
	return &KafkaBroker{
		producer: NewProducer(settings),
		consumer: NewConsumer(settings),
	}
}

func (b *KafkaBroker) Producer() common.Producer {
	return b.producer
}

func (b *KafkaBroker) Consumer() common.Consumer {
	return b.consumer
}

func (b *KafkaBroker) Close() error {
	return b.producer.Close()
}
