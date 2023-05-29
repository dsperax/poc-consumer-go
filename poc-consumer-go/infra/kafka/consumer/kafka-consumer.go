package consumer

import (
	"github.com/dsperax/poc-consumer-go/usecase"

	"github.com/segmentio/kafka-go"
)

type IKafkaConsumer interface {
	ConsumirMensagem()
}

type KafkaConsumerImpl struct {
}

var (
	canal        chan kafka.Message
	businesscase usecase.IKafkaConsumerUseCase
)

func NewKafkaConsumer(canalmsg chan kafka.Message, businesscaseinput usecase.IKafkaConsumerUseCase) IKafkaConsumer {
	canal = canalmsg
	businesscase = businesscaseinput
	return &KafkaConsumerImpl{}
}

func (*KafkaConsumerImpl) ConsumirMensagem() {
	for msg := range canal {
		businesscase.Executa(msg)
	}
}
