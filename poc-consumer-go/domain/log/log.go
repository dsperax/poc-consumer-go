package log

import (
	"encoding/json"
	"log"
	"time"
)

type Log struct {
	TsGravacao time.Time `json:"tsLog"`
	Conteudo   string    `json:"conteudo"`
}

func GravarLog(conteudolog any) {
	conteudo, err := json.Marshal(conteudolog)
	if err != nil {
		return
	}

	logprint := Log{
		TsGravacao: time.Now(),
		Conteudo:   string(conteudo),
	}

	logstring, err := json.Marshal(logprint)
	if err != nil {
		return
	}

	log.SetFlags(0)
	log.Println(string(logstring))
}
