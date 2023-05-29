package router

import (
	"fmt"
	"log"

	"github.com/dsperax/poc-consumer-go/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/sinhashubham95/go-actuator"
)

func Roteador() {
	port := utils.GetVariavelAmbiente("app_port", "8080")

	r := gin.Default()
	rotaActuator(r)
	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}

func rotaActuator(router *gin.Engine) {
	actuatorHandler := actuator.GetActuatorHandler(nil)
	ginActuatorHandler := func(c *gin.Context) {
		actuatorHandler(c.Writer, c.Request)
	}

	router.GET("/actuator/*endpoint", ginActuatorHandler)
}
