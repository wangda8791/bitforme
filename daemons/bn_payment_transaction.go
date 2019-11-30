package main

import (
	"time"

	"github.com/bn_funds/models"
	"github.com/bn_funds/models/worker"
	"github.com/bn_funds/services/wallet_service"
	"github.com/bn_funds/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

var (
	paymentTransaction worker.PaymentTransaction
	id                 string
)

func main() {
	models.Init_Logger("../log/bn_payment_transaction")
	utils.LoadEnvVars()
	gin.SetMode(utils.GetEnv("GIN_MODE", "release"))

	models.AMQPQueue_ = &models.AMQPQueue{}
	models.AMQPQueue_.Init()

	defer models.AMQPQueue_.Con.Close()
	defer models.AMQPQueue_.Ch.Close()

	models.Init(utils.GetEnv("DB_URL", ""))
	wallet_service.Init_Rpc()
	models.Init_State_Machine()

	paymentTransaction = worker.PaymentTransaction{}
	forever := make(chan bool)
	go func() {
		for {
			paymentTransaction.Process()
			time.Sleep(5 * time.Second)
		}
	}()

	logger.Infof(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
