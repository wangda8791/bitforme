// Package Monerod is client librari for Monerod JSON RPC API
package wallet_service

import (
	"strings"

	"github.com/bn_funds/structs"
	"github.com/bn_funds/utils"
	"github.com/google/logger"
)

// A Monerod represents a Monerod client
type Monerod struct {
	client       *rpcClient
	walletClient *rpcClient
	mainAddress  string
}

const (
	// DEFAULT_RPCCLIENT_TIMEOUT represent http timeout for rcp client
	ONE_XMR = 1000000000000.0 // 10^12
)

// New return a new Monerod
func NewMonerod(host string, port int, wPort int, user, passwd string, useSSL bool, mainAddress string) (*Monerod, error) {
	client, err := newClient(host, port, user, passwd, useSSL)
	client.serverAddr = client.serverAddr + "/json_rpc"
	walletClient, err := newClient(host, wPort, user, passwd, useSSL)
	walletClient.serverAddr = walletClient.serverAddr + "/json_rpc"
	if err != nil {
		return nil, err
	}
	return &Monerod{client, walletClient, mainAddress}, nil
}

// GetBalance return the balance of the server or of a specific account
//If [account] is "", returns the server's total available balance.
//If [account] is specified, returns the balance in the account
func (b *Monerod) GetBalance(account string, address string, minconf uint64) (balance float64, err error) {
	params := make(map[string]interface{}, 0)
	params["account_index"] = 0
	r, err := b.walletClient.call("getbalance", params)
	if err = handleError(err, &r); err != nil {
		return
	}
	var result map[string]float64
	utils.ByteArrayToJSON(r.Result, &result)

	balance = result["balance"] / ONE_XMR
	return
}

// GetNewAddress return a new address for account [account].
func (b *Monerod) GetNewAddress(account ...string) (addr string, err error) {

	addr = b.mainAddress + "-" + utils.RandHex(64)

	return
}

// GetTransaction returns a Monerod.Transation struct about the given transaction
func (b *Monerod) GetTransaction(txid string) (details []structs.TransactionDetails, err error) {
	params := make(map[string]string, 0)
	params["txid"] = txid
	r, err := b.walletClient.call("get_transfer_by_txid", params)

	if err = handleError(err, &r); err != nil {
		return
	}
	var raw map[string]interface{}
	utils.ByteArrayToJSON(r.Result, &raw)

	var rawDetails []map[string]interface{}

	utils.InterfaceToJSON(raw["transfers"], &rawDetails)

	details = make([]structs.TransactionDetails, len(rawDetails))
	j := 0
	for _, detail := range rawDetails {

		if detail["address"].(string) == b.mainAddress {
			var item structs.TransactionDetails
			item.Account = ""
			item.Address = detail["address"].(string) + "-" + detail["payment_id"].(string)
			item.Amount = detail["amount"].(float64) / ONE_XMR
			item.Fee = detail["fee"].(float64) / ONE_XMR
			item.Confirmations = int64(detail["confirmations"].(float64))
			item.TimeReceived = int64(detail["timestamp"].(float64))
			if detail["type"].(string) == "in" {
				item.Category = "receive"
			} else {
				item.Category = "send"
			}
			details[j] = item
			j += 1
		}
	}

	return
}

// SendToAddress send an amount to a given address
func (b *Monerod) SendToAddress(fromAddress string, toAddress string, amount float64, comment, commentTo string, bInteranl bool) (txID string, err error) {

	params := make(map[string]interface{})
	destinations := make([]map[string]interface{}, 1)
	destination := make(map[string]interface{})
	destination["amount"] = int64(amount * ONE_XMR)
	s := strings.Split(toAddress, "-")
	destination["address"] = s[0]
	if len(s) > 1 {
		params["payment_id"] = s[1]
	}
	destinations[0] = destination
	params["destinations"] = destinations

	params["account_index"] = 0
	params["priority"] = 0
	params["ring_size"] = 11
	params["get_tx_key"] = true
	params["subaddr_indices"] = []interface{}{0}

	r, err := b.walletClient.call("transfer", params)
	logger.Info(string(r.Result), err)
	if err = handleError(err, &r); err != nil {
		return
	}
	var result map[string]interface{}
	utils.ByteArrayToJSON(r.Result, &result)
	txID = result["tx_hash"].(string)
	return
}

func (b *Monerod) GetDepositTransactions(last int) (txids []string, latestBlockNumber int64, err error) {
	params := make(map[string]interface{})
	params["account_index"] = 0
	params["in"] = true
	params["filter_by_height"] = true
	params["min_height"] = last
	blockNumber, err := b.GetBlockNumber()
	params["max_height"] = blockNumber + 1

	r, err := b.walletClient.call("get_transfers", params)
	if err = handleError(err, &r); err != nil {
		return
	}

	var raw map[string]interface{}
	utils.ByteArrayToJSON(r.Result, &raw)

	var rawTransactions []map[string]interface{}

	utils.InterfaceToJSON(raw["in"], &rawTransactions)
	txids = make([]string, 0)

	for _, rawTransaction := range rawTransactions {
		var transaction map[string]interface{}
		utils.InterfaceToJSON(rawTransaction, &transaction)
		txids = append(txids, transaction["txid"].(string))
	}
	latestBlockNumber = blockNumber
	return
}
func (b *Monerod) GetBlockNumber() (blockNumber int64, err error) {

	r, err := b.walletClient.call("getheight", []interface{}{})
	if err = handleError(err, &r); err != nil {
		return
	}
	var raw map[string]interface{}
	utils.ByteArrayToJSON(r.Result, &raw)
	blockNumber = int64(raw["height"].(float64))
	return
}

// GetNewAddress return a new address for account [account].
func (b *Monerod) GetMainAddress() (addr string) {
	return b.mainAddress
}
