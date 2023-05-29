package usecase

import (
	"encoding/json"
	"fmt"

	domaindatabase "github.com/dsperax/poc-consumer-go/domain/database"
	domain "github.com/dsperax/poc-consumer-go/domain/kafka"
	"github.com/dsperax/poc-consumer-go/domain/log"
	"github.com/dsperax/poc-consumer-go/infra/repository"

	"github.com/segmentio/kafka-go"
)

type BusinessUseCase struct {
}

var (
	canalresultado chan bool
	businessrepo   repository.IBusinessRepository
)

func NewBusinessUseCase(repository repository.IBusinessRepository, canalresult chan bool) IKafkaConsumerUseCase {
	canalresultado = canalresult
	businessrepo = repository
	return &BusinessUseCase{}
}

func (s *BusinessUseCase) Executa(mensagemkafka kafka.Message) {
	businesskafkadomain, businessentitydomain, err := s.geraEntidade(mensagemkafka)
	if err != nil {
		log.GravarLog(err.Error())
		canalresultado <- true
		return
	}

	if businesskafkadomain.Status == "I" {
		err = businessrepo.Excluir(businessentitydomain)
	} else {
		existe, err := businessrepo.ExisteRegistro(businessentitydomain)
		if err != nil {
			canalresultado <- false
			return
		}

		if existe {
			err = businessrepo.Atualizar(businessentitydomain)
		} else {
			err = businessrepo.Inserir(businessentitydomain)
		}
	}

	log.GravarLog(fmt.Sprintf("Mensagem processada: Op[%s] %s %s",
		businesskafkadomain.CodigoCliente,
		businesskafkadomain.NomeCliente,
		businesskafkadomain.Regiao))
	canalresultado <- err == nil
}

func (r *BusinessUseCase) geraEntidade(mensagemkafka kafka.Message) (domain.BusinessKafka, domaindatabase.BusinessEntity, error) {
	var businessentitydomain domaindatabase.BusinessEntity
	var businesskafkadomain domain.BusinessKafka
	err := json.Unmarshal(mensagemkafka.Value, &businesskafkadomain)
	if err != nil {
		log.GravarLog(err.Error())
		return businesskafkadomain, businessentitydomain, err
	}

	err = businessentitydomain.FromDTO(businesskafkadomain)
	if err != nil {
		log.GravarLog(err.Error())
		return businesskafkadomain, businessentitydomain, err
	}

	return businesskafkadomain, businessentitydomain, nil
}
