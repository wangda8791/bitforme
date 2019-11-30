package wallet_service

import "github.com/bn_funds/structs"

type Coind interface {
	GetTransaction(txid string) ([]structs.TransactionDetails, error)
	GetMainAddress() string
	GetNewAddress(...string) (string, error)
	GetBalance(account string, address string, minconf uint64) (balance float64, err error)
	SendToAddress(fromAddress string, toAddress string, amount float64, comment, commentTo string, bInternal bool) (txID string, err error)
	GetDepositTransactions(last int) (txids []string, latestBlockNumber int64, err error)
	GetBlockNumber() (latestBlockNumber int64, err error)
}
