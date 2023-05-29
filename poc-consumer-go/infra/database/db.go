package database

import (
	"fmt"
	"os"

	"github.com/dsperax/poc-consumer-go/app/utils"
	"github.com/dsperax/poc-consumer-go/domain/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := utils.GetVariavelAmbiente("PG_POC_CONSUMER_GO", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", "localhost", "postgres", "root", "postgres", 5432, "disable"))
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.GravarLog("Erro ao conectar com o banco de dados!")
		os.Exit(1)
	}
}
