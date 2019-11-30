// Package Nemd is client librari for Nemd JSON RPC API
package wallet_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/bn_funds/structs"
	"github.com/bn_funds/utils"
	"github.com/google/logger"
)

// A Nemd represents a Stellard client
type Nemd struct {
	serverAddr  string
	apiKey      string
	mainAddress string
}

const (
// DEFAULT_RPCCLIENT_TIMEOUT represent http timeout for rcp client
// ONE_WAVES = 100000000.0
)

// New return a new Stellard
func NewNemd(host string, port int, useSSL bool, apiKey string, mainAddress string) (*Nemd, error) {

	var serverAddr string
	if useSSL {
		serverAddr = "https://"
	} else {
		serverAddr = "http://"
	}
	return &Nemd{fmt.Sprintf("%s%s:%d", serverAddr, host, port), apiKey, mainAddress}, nil
}
func (self *Nemd) client() *http.Client {
	client := &http.Client{}
	return client
}
func (self *Nemd) request(method, path string, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", self.serverAddr, path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	if self.apiKey != "" {
		req.Header.Add("X-API-KEY", self.apiKey)
	}
	return req, nil
}

func (self *Nemd) Get(path string) (out []byte, err error) {
	req, err := self.request("GET", path, nil)

	if err != nil {
		return
	}

	res, err := self.client().Do(req)

	if err != nil {
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	out = data
	return
}

// doTimeoutRequest process a HTTP request with timeout
func (self *Nemd) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := self.client().Do(req)
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("Timeout reading data from server")
	}
}
func (self *Nemd) PostBody(path string, body interface{}) (out []byte, err error) {

	payloadBuffer := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(payloadBuffer)
	err = jsonEncoder.Encode(body)
	if err != nil {
		return
	}
	req, err := self.request("POST", path, payloadBuffer)
	logger.Info("PostBody:", req, err)

	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// res, err := self.client().Do(req)
	connectTimer := time.NewTimer(RPCCLIENT_TIMEOUT * time.Second)

	res, err := self.doTimeoutRequest(connectTimer, req)

	if err != nil {
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	out = data

	return
}

// GetBalance return the balance of the server or of a specific account
//If [account] is "", returns the server's total available balance.
//If [account] is specified, returns the balance in the account
func (self *Nemd) GetBalance(account string, address string, minconf uint64) (balance float64, err error) {

	out, err := self.Get("addresses/balance/details/" + address)
	var raw map[string]interface{}
	utils.ByteArrayToJSON(out, &raw)
	balance = raw["available"].(float64) / ONE_WAVES

	return
}

// GetNewAddress return a new address for account [account].
func (self *Nemd) GetNewAddress(account ...string) (addr string, err error) {

	addr = self.mainAddress + "-" + utils.RandSeq(10)

	return
}

func (self *Nemd) GetAccounts() (accounts []string, err error) {

	out, err := self.Get("addresses")
	json.Unmarshal(out, &accounts)
	return
}

func (self *Nemd) GetBlockNumber() (blockNumber int64, err error) {
	out, err := self.Get("blocks/headers/last")
	var raw map[string]interface{}
	utils.ByteArrayToJSON(out, &raw)
	blockNumber = int64(raw["height"].(float64))
	return
}

// GetTransaction returns a Nemd.Transation struct about the given transaction
func (self *Nemd) GetTransaction(txid string) (details []structs.TransactionDetails, err error) {
	out, err := self.Get("transactions/info/" + txid)
	var raw map[string]interface{}
	utils.ByteArrayToJSON(out, &raw)

	details = make([]structs.TransactionDetails, 1)

	var item structs.TransactionDetails
	item.Account = ""
	accounts, err := self.GetAccounts()
	var i int
	var account string
	for i, account = range accounts {

		if account == raw["sender"].(string) {
			item.Address = raw["sender"].(string)
			item.Category = "send"
			break
		}

		if account == raw["recipient"].(string) {
			item.Address = raw["recipient"].(string)
			item.Category = "receive"
			break
		}
	}

	if i >= len(accounts) {
		err = errors.New("This Transaction is not related our wallet.")
		return
	}
	item.Amount = raw["amount"].(float64) / ONE_WAVES
	item.Fee = raw["fee"].(float64) / ONE_WAVES

	latestBlockNumber, err := self.GetBlockNumber()
	currentBlockNumber := int64(raw["height"].(float64))
	item.Confirmations = latestBlockNumber - currentBlockNumber
	item.TimeReceived = int64(raw["timestamp"].(float64)) / 1000
	details[0] = item

	return
}

// GetNewAddress return a new address for account [account].
func (self *Nemd) GetMainAddress() (addr string) {
	return self.mainAddress
}

// SendToAddress send an amount to a given address
func (self *Nemd) SendToAddress(fromAddress string, toAddress string, amount float64, comment, commentTo string, bInternal bool) (txID string, err error) {

	if fromAddress == "" {
		fromAddress = self.mainAddress
	}
	params := make(map[string]interface{}, 0)
	// params["assetId"] = ""
	params["amount"] = amount * ONE_WAVES
	params["fee"] = 100000
	params["sender"] = fromAddress
	params["recipient"] = toAddress
	// params["attachment"] = comment

	if bInternal == true {
		params["amount"] = amount*ONE_WAVES - 100000
	}

	logger.Info("Params", params)

	out, err := self.PostBody("assets/transfer", params)
	if err != nil {
		return
	}
	var raw map[string]interface{}
	utils.ByteArrayToJSON(out, &raw)
	logger.Info("SendToAddress:", raw)
	if raw["error"] != nil {
		err = errors.New(raw["message"].(string))
	} else {
		txID = raw["id"].(string)
	}
	return
}

func (self *Nemd) GetDepositTransactions(last int) (txids []string, latestBlockNumber int64, err error) {
	// last = 404267
	// latestBlockNumber = int64(last) + 5
	// logger.Info("Block:", last)
	latestBlockNumber, err = self.GetBlockNumber()
	// logger.Info("BlockNumber", blockNumber)
	if err != nil {
		return
	}

	txids = make([]string, 0)
	for i := last; i < int(latestBlockNumber); i++ {
		txidsperblock, _ := self.GetDepositTransactionsByBlock(i)
		txids = append(txids, txidsperblock...)
	}
	if len(txids) != 0 {
		logger.Info("Txids:", txids)
	}
	return
}

func (self *Nemd) GetDepositTransactionsByBlock(blockNumber int) (txids []string, err error) {
	out, err := self.Get("blocks/at/" + strconv.Itoa(blockNumber))

	txids = make([]string, 0)
	accounts, err := self.GetAccounts()
	var block map[string]interface{}
	utils.ByteArrayToJSON(out, &block)

	var transactions []map[string]interface{}
	utils.InterfaceToJSON(block["transactions"], &transactions)
	for _, transaction := range transactions {
		// logger.Info("Transaction:", transaction)
		if transaction["sender"] == self.mainAddress {
			continue
		}
		for _, account := range accounts {
			if account == self.mainAddress {
				continue
			}
			if transaction["recipient"] == account {
				txids = append(txids, transaction["id"].(string))
			}
		}
	}

	return
}
