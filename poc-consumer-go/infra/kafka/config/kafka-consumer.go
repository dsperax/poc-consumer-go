package akafka

import (
	"context"
	"fmt"
	"time"

	"github.com/dsperax/poc-consumer-go/domain/log"

	"github.com/segmentio/kafka-go"
)

func Consumer(topic string, servers []string, groupid string, mensagem chan kafka.Message, resultado chan bool) {
	kafkaConsumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           servers,
		Topic:             topic,
		GroupID:           groupid,
		HeartbeatInterval: time.Second * 30,
		// Logger:            kafka.LoggerFunc(logf),
		// ErrorLogger:       kafka.LoggerFunc(logf),
	})

	ctx := context.Background()

	for {
		msg, err := kafkaConsumer.FetchMessage(ctx)
		if err != nil {
			log.GravarLog(fmt.Sprintf("Erro: %s", err.Error()))
			break
		}

		mensagem <- msg

		if <-resultado {
			kafkaConsumer.CommitMessages(ctx, msg)
		} else {
			continue
		}
	}
}
