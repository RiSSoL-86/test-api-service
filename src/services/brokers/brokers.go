package brokers

import (
	brokersettings "app/src/app_settings/brokers"
	"app/src/core/brokers/common"
	"app/src/services/brokers/kafka"
)

func NewKafkaBroker(settings *brokersettings.KafkaSettings) common.Broker {
	return kafka.NewBroker(settings)
}
