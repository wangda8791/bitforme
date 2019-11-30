package main

import (
	"github.com/bn_funds/models"
	"github.com/bn_funds/routers"
	"github.com/bn_funds/services/wallet_service"
	"github.com/bn_funds/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	models.Init_Logger("../log/bn_main")
	utils.LoadEnvVars()
	gin.SetMode(utils.GetEnv("GIN_MODE", "release"))
	models.Init_Cred()
	// models.Init_Pusher()
	// models.Init_Redis()
	models.AMQPQueue_ = &models.AMQPQueue{}
	models.AMQPQueue_.Init()

	defer models.AMQPQueue_.Con.Close()
	defer models.AMQPQueue_.Ch.Close()

	models.Init(utils.GetEnv("DB_URL", ""))

	wallet_service.Init_Rpc()
	models.Init_State_Machine()
	routers.Init()
}
