package repository

import (
	"fmt"

	domain "github.com/dsperax/poc-consumer-go/domain/database"
	"github.com/dsperax/poc-consumer-go/domain/log"
	"github.com/dsperax/poc-consumer-go/infra/database"
)

const (
	filtrochave = "CODIGO_CLIENTE=? and FILIAL=? and REGIAO=?"
)

type BusinessRepoPg struct {
}

func NewBusinessMSSQLRepository() IBusinessRepository {
	return &BusinessRepoPg{}
}

func (r *BusinessRepoPg) Inserir(businessdomain domain.BusinessEntity) error {
	err := database.DB.Create(&businessdomain).Error
	if err != nil {
		log.GravarLog(err.Error())
		return err
	}

	return nil
}

func (r *BusinessRepoPg) Excluir(businessdomain domain.BusinessEntity) error {
	err := database.DB.Where(
		filtrochave,
		businessdomain.CodigoCliente,
		businessdomain.Filial,
		businessdomain.Regiao).Delete(&businessdomain).Error
	if err != nil {
		log.GravarLog(err.Error())
		return err
	}

	return nil
}

func (r *BusinessRepoPg) ExisteRegistro(businessdomain domain.BusinessEntity) (bool, error) {
	err := database.DB.Where(
		filtrochave,
		businessdomain.CodigoCliente,
		businessdomain.Filial,
		businessdomain.Regiao).First(&businessdomain).Error
	if err != nil && err.Error() != "record not found" {
		log.GravarLog(fmt.Sprintf("Deu erro no find: %s", err.Error()))
		return false, err
	}

	if err != nil && err.Error() == "record not found" {
		return false, nil
	}
	
	return true, nil
}

func (r *BusinessRepoPg) Atualizar(businessdomain domain.BusinessEntity) error {
	err := database.DB.Where(
		filtrochave,
		businessdomain.CodigoCliente,
		businessdomain.Filial,
		businessdomain.Regiao).Save(&businessdomain).Error
	if err != nil {
		log.GravarLog(err.Error())
		return err
	}

	return nil
}
