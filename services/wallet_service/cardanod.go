// Package Cardanod is client librari for Cardanod JSON RPC API
package wallet_service

import (
	"bytes"
	"crypto/tls"
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

// A Cardanod represents a Stellard client
type Cardanod struct {
	serverAddr   string
	mainAddress  string
	spendingPass string
	accountIndex int64
}

const (
	// DEFAULT_RPCCLIENT_TIMEOUT represent http timeout for rcp client
	ONE_ADA = 1000000.0
)

// New return a new Stellard
func NewCardanod(host string, port int, useSSL bool, mainAddress string, spendingPass string, accountIndex string) (*Cardanod, error) {

	var serverAddr string
	if useSSL {
		serverAddr = "https://"
	} else {
		serverAddr = "http://"
	}
	index, _ := strconv.ParseInt(accountIndex, 10, 64)
	return &Cardanod{fmt.Sprintf("%s%s:%d", serverAddr, host, port), mainAddress, spendingPass, index}, nil
}
func (self *Cardanod) client() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client
}
func (self *Cardanod) request(method, path string, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", self.serverAddr, path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json;charset=utf-8")

	return req, nil
}

func (self *Cardanod) Get(path string) (out []byte, err error) {
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
func (self *Cardanod) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
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
func (self *Cardanod) PostBody(path string, body interface{}) (out []byte, err error) {

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
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Accept", "application/json;charset=utf-8")

	// res, err := self.client().Do(req)
	connectTimer := time.NewTimer(RPCCLIENT_TIMEOUT * time.Second)

	res, err := self.doTimeoutRequest(connectTimer, req)
	logger.Info("Res:", res, err)
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
func (self *Cardanod) GetBalance(account string, address string, minconf uint64) (balance float64, err error) {

	out, err := self.Get("api/v1/wallets/" + self.GetMainAddress() + "/accounts/" + utils.I64tos(self.accountIndex))
	var raw map[string]interface{}
	utils.ByteArrayToJSON(out, &raw)
	var data map[string]interface{}
	utils.InterfaceToJSON(raw["data"], &data)

	balance = data["amount"].(float64) / ONE_ADA

	return
}

// GetNewAddress return a new address for account [account].
func (self *Cardanod) GetNewAddress(account ...string) (addr string, err error) {

	params := make(map[string]interface{}, 0)
	// params["assetId"] = ""
	params["accountIndex"] = self.accountIndex
	params["walletId"] = self.GetMainAddress()
	params["spendingPassword"] = self.spendingPass
	out, err := self.PostBody("api/v1/addresses", params)

	var raw map[string]interface{}
	utils.ByteArrayToJSON(out, &raw)
	var data map[string]interface{}
	utils.InterfaceToJSON(raw["data"], &data)
	addr = data["id"].(string)

	return
}

func (self *Cardanod) GetAccounts() (accounts []string, err error) {

	out, err := self.Get("addresses")
	json.Unmarshal(out, &accounts)
	return
}

func (self *Cardanod) GetBlockNumber() (blockNumber int64, err error) {

	blockNumber = utils.GetCurrentTimeStamp()
	return
}

// GetTransaction returns a Cardanod.Transation struct about the given transaction
func (self *Cardanod) GetTransaction(txid string) (details []structs.TransactionDetails, err error) {

	out, err := self.Get("api/v1/transactions/?" + "wallet_id=" + self.mainAddress + "&" + "id=" + txid)
	var raw map[string]interface{}
	utils.ByteArrayToJSON(out, &raw)
	var data []map[string]interface{}
	utils.InterfaceToJSON(raw["data"], &data)
	details = make([]structs.TransactionDetails, len(data))
	j := 0
	for _, dataItem := range data {
		var item structs.TransactionDetails
		if dataItem["direction"].(string) == "incoming" {
			item.Category = "receive"
		} else if dataItem["direction"].(string) == "outgoing" {
			item.Category = "send"
		}

		var outputs []map[string]interface{}
		utils.InterfaceToJSON(dataItem["outputs"], &outputs)
		if len(outputs) < 2 {
			continue
		}

		item.Address = outputs[1]["address"].(string)
		item.Amount = outputs[1]["amount"].(float64) / ONE_ADA
		item.Account = ""
		item.Fee = 0
		item.Confirmations = int64(dataItem["confirmations"].(float64))
		item.TimeReceived = 0 // temp
		details[j] = item
		j += 1
	}

	return
}

// GetNewAddress return a new address for account [account].
func (self *Cardanod) GetMainAddress() (addr string) {
	return self.mainAddress
}

// SendToAddress send an amount to a given address
func (self *Cardanod) SendToAddress(fromAddress string, toAddress string, amount float64, comment, commentTo string, bInternal bool) (txID string, err error) {

	if fromAddress == "" {
		fromAddress = self.mainAddress
	}
	params := make(map[string]interface{}, 0)
	// params["assetId"] = ""
	destinations := make([]map[string]interface{}, 1)
	destination := make(map[string]interface{})
	destination["amount"] = amount * ONE_ADA
	destination["address"] = toAddress
	destinations[0] = destination
	params["destinations"] = destinations
	source := make(map[string]interface{})
	source["accountIndex"] = self.accountIndex
	source["walletId"] = self.GetMainAddress()
	params["source"] = source
	params["spendingPassword"] = self.spendingPass

	logger.Info("Params", params)

	out, err := self.PostBody("api/v1/transactions", params)
	if err != nil {
		return
	}
	var raw map[string]interface{}
	utils.ByteArrayToJSON(out, &raw)
	logger.Info("SendToAddress:", raw)
	var data map[string]interface{}
	utils.InterfaceToJSON(raw["data"], &data)

	// for _, item := range data {
	txID = data["id"].(string)
	// break
	// }

	return
}

func (self *Cardanod) GetDepositTransactions(last int) (txids []string, latestBlockNumber int64, err error) {
	latestBlockNumber, err = self.GetBlockNumber()
	if err != nil {
		return
	}

	out, err := self.Get("api/v1/transactions?" + "wallet_id=" + self.GetMainAddress() + "&account_index=" + utils.I64tos(self.accountIndex) + "&created_at=" + "RANGE[" + utils.I64tos(int64(last)) + "," + utils.I64tos(latestBlockNumber) + "]")

	txids = make([]string, 0)
	var raw map[string]interface{}
	utils.ByteArrayToJSON(out, &raw)

	var transactions []map[string]interface{}
	utils.InterfaceToJSON(raw["data"], &transactions)
	for _, transaction := range transactions {
		// logger.Info("Transaction:", transaction)
		details, err := self.GetTransaction(transaction["id"].(string))
		if err != nil {
			continue
		}
		for _, detail := range details {
			if detail.Category == "receive" {
				txids = append(txids, transaction["id"].(string))
			}
		}
	}
	latestBlockNumber += 1
	return
}

func (self *Cardanod) GetDepositTransactionsByBlock(blockNumber int) (txids []string, err error) {
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
