package main

import (
	"github.com/bn_funds/models"
	"github.com/bn_funds/routers"
	"github.com/bn_funds/services/wallet_service"
	"github.com/bn_funds/utils"
	"github.com/gin-gonic/gin"
)

var (
// deposit_coin worker.DepositCoin
// deposit_coin_address worker.DepositCoinAddress
)

func main() {
	models.Init_Logger("../log/bn_deposit_coin")
	utils.LoadEnvVars()
	gin.SetMode(utils.GetEnv("GIN_MODE", "release"))

	models.Init(utils.GetEnv("DB_URL", ""))

	models.AMQPQueue_ = &models.AMQPQueue{}
	models.AMQPQueue_.Init()

	defer models.AMQPQueue_.Con.Close()
	defer models.AMQPQueue_.Ch.Close()

	wallet_service.Init_Rpc()
	models.Init_State_Machine()

	routers.Init_Wallet()
}
