package usecase

import (
	"github.com/segmentio/kafka-go"
)

type IKafkaConsumerUseCase interface {
	Executa(mensagem kafka.Message)
}
