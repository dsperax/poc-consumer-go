package repository

import (
	domain "github.com/dsperax/poc-consumer-go/domain/database"
)

type IBusinessRepository interface {
	ExisteRegistro(businessdomain domain.BusinessEntity) (bool, error)
	Inserir(businessdomain domain.BusinessEntity) error
	Atualizar(businessdomain domain.BusinessEntity) error
	Excluir(businessdomain domain.BusinessEntity) error
}
