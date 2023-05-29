package domain

import (
	"strconv"
	"time"

	domainkafka "github.com/dsperax/poc-consumer-go/domain/kafka"
)

type BusinessEntity struct {
	CodigoCliente int       `gorm:"column:codigo_cliente"`
	NomeCliente   string    `gorm:"column:nome_cliente"`
	Telefone      string    `gorm:"column:telefone"`
	DataCadastro  time.Time `gorm:"column:data_cadastro"`
	Filial        int       `gorm:"column:filial"`
	Regiao        int       `gorm:"column:regiao"`
	Status        string    `gorm:"column:status"`
}

func (*BusinessEntity) TableName() string {
	return "clientesvendas"
}

func (p *BusinessEntity) FromDTO(mensagemkafka domainkafka.BusinessKafka) error {
	var err error = nil

	p.NomeCliente = mensagemkafka.NomeCliente

	p.CodigoCliente, err = strconv.Atoi(mensagemkafka.CodigoCliente)
	if err != nil {
		return err
	}

	p.Telefone = mensagemkafka.Telefone

	p.DataCadastro, err = time.Parse("2006-01-02", mensagemkafka.DataCadastro)
	if err != nil {
		return err
	}

	p.Filial, err = strconv.Atoi(mensagemkafka.Filial)
	if err != nil {
		return err
	}

	p.Regiao, err = strconv.Atoi(mensagemkafka.Regiao)
	if err != nil {
		return err
	}

	p.Status = mensagemkafka.Status

	return nil
}
