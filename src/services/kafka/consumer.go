package kafka

import (
	brokersettings "app/src/app_settings/brokers"
	"app/src/core/brokers/common"
	"context"

	kgo "github.com/segmentio/kafka-go"
)

type Consumer struct {
	brokerAddresses []string
}

func NewConsumer(settings *brokersettings.KafkaSettings) *Consumer {
	return &Consumer{brokerAddresses: settings.BrokerAddresses}
}

func (c *Consumer) Consume(ctx context.Context, topic string, groupID string, handler common.MessageHandler) error {
	reader := kgo.NewReader(kgo.ReaderConfig{
		Brokers: c.brokerAddresses,
		Topic:   topic,
		GroupID: groupID,
	})
	defer reader.Close()

	for {
		kafkaMessage, err := reader.FetchMessage(ctx)
		if err != nil {
			return err
		}

		headers := make(map[string]string, len(kafkaMessage.Headers))
		for _, header := range kafkaMessage.Headers {
			headers[header.Key] = string(header.Value)
		}

		message := common.Message{
			Topic:   kafkaMessage.Topic,
			Key:     kafkaMessage.Key,
			Value:   kafkaMessage.Value,
			Headers: headers,
		}

		if err := handler(ctx, message); err != nil {
			return err
		}

		if err := reader.CommitMessages(ctx, kafkaMessage); err != nil {
			return err
		}
	}
}
