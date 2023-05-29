package main

import (
	"strings"

	"github.com/dsperax/poc-consumer-go/app/router"
	"github.com/dsperax/poc-consumer-go/app/utils"
	"github.com/dsperax/poc-consumer-go/infra/database"
	akafka "github.com/dsperax/poc-consumer-go/infra/kafka/config"
	"github.com/dsperax/poc-consumer-go/infra/kafka/consumer"
	"github.com/dsperax/poc-consumer-go/infra/repository"
	"github.com/dsperax/poc-consumer-go/usecase"

	"github.com/segmentio/kafka-go"
)

var (
	canal           chan kafka.Message             = make(chan kafka.Message)
	canalresultado  chan bool                      = make(chan bool)
	businessrepo    repository.IBusinessRepository = repository.NewBusinessMSSQLRepository()
	businessusecase usecase.IKafkaConsumerUseCase  = usecase.NewBusinessUseCase(businessrepo, canalresultado)
	consumerimpl    consumer.IKafkaConsumer        = consumer.NewKafkaConsumer(canal, businessusecase)
)

func main() {
	database.ConectaComBancoDeDados()

	topico := utils.GetVariavelAmbiente("kafka_topic", "topico-vendas")
	servidores := strings.Split(utils.GetVariavelAmbiente("kafka_servers", "localhost:9092"), ";")
	groupid := utils.GetVariavelAmbiente("kafka_groupid", "consumer-poc-go")

	go akafka.Consumer(topico, servidores, groupid, canal, canalresultado)
	go consumerimpl.ConsumirMensagem()

	router.Roteador()
}
