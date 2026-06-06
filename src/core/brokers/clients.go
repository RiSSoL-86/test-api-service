package brokers

import (
	brokersettings "app/src/app_settings/brokers"
	"app/src/core/brokers/common"
	"app/src/services/kafka"
)

func NewKafkaBroker(settings *brokersettings.KafkaSettings) common.Broker {
	return kafka.NewKafkaBroker(settings)
}
