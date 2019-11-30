package api_v1

import (
	"net/http"

	"github.com/bn_funds/models/worker"
	"github.com/gin-gonic/gin"
)

func Wallet_Notify(c *gin.Context) {
	currency := c.Request.URL.Query().Get("currency")
	txid := c.Request.URL.Query().Get("txid")
	(&worker.DepositCoin{}).Process(currency, txid)

	c.JSON(http.StatusOK, "")
}
