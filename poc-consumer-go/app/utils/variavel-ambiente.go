package utils

import "os"

func GetVariavelAmbiente(nome string, padrao string) string {
	if valor, ok := os.LookupEnv(nome); ok {
		return valor
	} else {
		return padrao
	}
}
