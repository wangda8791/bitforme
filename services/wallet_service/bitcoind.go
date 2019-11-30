// Package Bitcoind is client librari for bitcoind JSON RPC API
package wallet_service

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/bn_funds/structs"
	"github.com/bn_funds/utils"
)

// A Bitcoind represents a Bitcoind client
type Bitcoind struct {
	client *rpcClient
}

// New return a new bitcoind
func NewBitcoind(host string, port int, user, passwd string, useSSL bool) (*Bitcoind, error) {
	rpcClient, err := newClient(host, port, user, passwd, useSSL)
	if err != nil {
		return nil, err
	}
	return &Bitcoind{rpcClient}, nil
}

// GetBalance return the balance of the server or of a specific account
//If [account] is "", returns the server's total available balance.
//If [account] is specified, returns the balance in the account
func (b *Bitcoind) GetBalance(account string, address string, minconf uint64) (balance float64, err error) {
	r, err := b.client.call("getbalance", []interface{}{account, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	balance, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

// GetNewAddress return a new address for account [account].
func (b *Bitcoind) GetNewAddress(account ...string) (addr string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("Bad parameters for GetNewAddress: you can set 0 or 1 account")
		return
	}
	r, err := b.client.call("getnewaddress", account)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addr)
	return
}

// GetTransaction returns a Bitcoind.Transation struct about the given transaction
func (b *Bitcoind) GetTransaction(txid string) (details []structs.TransactionDetails, err error) {
	r, err := b.client.call("gettransaction", []interface{}{txid})
	if err = handleError(err, &r); err != nil {
		return
	}

	var raw map[string]interface{}
	utils.ByteArrayToJSON(r.Result, &raw)

	var rawDetails []map[string]interface{}

	utils.InterfaceToJSON(raw["details"], &rawDetails)

	details = make([]structs.TransactionDetails, len(rawDetails))
	for i, detail := range rawDetails {
		var item structs.TransactionDetails
		item.Account = detail["account"].(string)
		if detail["address"] != nil {
			item.Address = detail["address"].(string)
		} else {
			item.Address = ""
		}
		item.Address = detail["address"].(string)
		item.Category = detail["category"].(string)
		item.Amount = detail["amount"].(float64)
		if detail["fee"] != nil {
			item.Fee = detail["fee"].(float64)
		} else {
			item.Fee = 0
		}
		item.Confirmations = int64(raw["confirmations"].(float64))
		item.TimeReceived = int64(raw["timereceived"].(float64))
		details[i] = item
	}

	return
}

// SendToAddress send an amount to a given address
func (b *Bitcoind) SendToAddress(fromAddress string, toAddress string, amount float64, comment, commentTo string, bInteranl bool) (txID string, err error) {
	r, err := b.client.call("sendtoaddress", []interface{}{toAddress, amount, comment, commentTo})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}

func (b *Bitcoind) GetDepositTransactions(last int) (txids []string, latestBlockNumber int64, err error) {

	return
}

// GetNewAddress return a new address for account [account].
func (b *Bitcoind) GetMainAddress() (addr string) {
	return ""
}

func (b *Bitcoind) GetBlockNumber() (latestBlockNumber int64, err error) {
	return
}
