package api_v1

import (
	"net/http"
	"strconv"

	"github.com/bn_funds/services/wallet_service"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

func Generate_Address(c *gin.Context) {

	currency := c.Request.URL.Query().Get("currency")
	var (
		addr string
		err  error
	)

	addr, err = wallet_service.CoinRPC[currency].(wallet_service.Coind).GetNewAddress("bitforme")

	res := make(map[string]interface{}, 0)
	if err == nil && addr != "" {
		res["success"] = true
		result := make(map[string]interface{})
		result["address"] = addr
		res["result"] = result
	} else {
		res["success"] = false
		error_ := make(map[string]interface{})
		error_["err_code"] = 500
		res["error"] = error_
		logger.Error("Failed to Generate Address: %+v", err)
	}

	c.JSON(http.StatusOK, res)
}

func Get_Transaction(c *gin.Context) {

	currency := c.Request.URL.Query().Get("currency")
	txid := c.Request.URL.Query().Get("txid")
	res := make(map[string]interface{}, 0)
	logger.Info("Currency, Txid", currency, txid)
	details, err := wallet_service.CoinRPC[currency].(wallet_service.Coind).GetTransaction(txid)
	if err != nil {
		logger.Info(err)
		res["success"] = false
		error_ := make(map[string]interface{})
		error_["err_code"] = 500
		error_["err_msg"] = err.Error()
		res["error"] = error_
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res["sucess"] = true
	result := make([]map[string]interface{}, len(details))
	for i, detail := range details {
		item := make(map[string]interface{})
		item["account"] = detail.Account
		item["address"] = detail.Address
		item["amount"] = detail.Amount
		item["category"] = detail.Category
		item["fee"] = detail.Fee
		item["confirmations"] = detail.Confirmations
		item["timereceived"] = detail.TimeReceived
		result[i] = item
	}
	res["result"] = result

	c.JSON(http.StatusOK, res)
}

func Withdraw(c *gin.Context) {
	currency := c.PostForm("currency")
	address := c.PostForm("address")

	var txid string

	res := make(map[string]interface{}, 0)

	amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)

	if err != nil {
		error_ := make(map[string]interface{})
		res["success"] = false
		error_["err_code"] = 400
		error_["err_msg"] = "Amount you inputed is not correct"
		res["error"] = error_
		c.JSON(http.StatusBadRequest, res)
		return
	}

	balance, _ := wallet_service.CoinRPC[currency].(wallet_service.Coind).GetBalance("bitforme", address, 0)
	logger.Info("balance:", balance)
	if balance >= amount {
		tx_id, err := wallet_service.CoinRPC[currency].(wallet_service.Coind).SendToAddress("", address, amount, "", "", false)
		txid = tx_id
		if err != nil {
			res["success"] = false
			error_ := make(map[string]interface{})
			error_["err_code"] = 500
			error_["err_msg"] = err.Error()
			res["error"] = error_
			c.JSON(http.StatusInternalServerError, res)
			return
		}

	} else {
		res["success"] = false
		error_ := make(map[string]interface{})
		error_["err_code"] = 400
		error_["err_msg"] = "Insufficient Balance"
		res["error"] = error_
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res["success"] = true
	result := make(map[string]interface{})
	result["txid"] = txid
	res["result"] = result

	c.JSON(http.StatusOK, res)
}
