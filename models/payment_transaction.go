package models

import (
	"time"

	"github.com/bn_funds/jsons"
	"github.com/bn_funds/services/wallet_service"
	"github.com/bn_funds/utils"
	"github.com/google/logger"
	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
	"github.com/shopspring/decimal"
)

type PaymentTransaction struct {
	gorm.Model

	Txid          string
	Amount        decimal.Decimal
	Confirmations int
	Address       string

	Receive_At time.Time
	Currency   int
	Txout      int

	transition.Transition

	// Deposit Deposit
}

// min_confirmation = 1, max_confirmation = 5
func (self *PaymentTransaction) Init(txid string, txout int, address string, amount decimal.Decimal,
	confirmations int, receive_at time.Time, currency int) {
	self.Txid = txid
	self.Txout = txout
	self.Address = address
	self.Amount = amount
	self.Confirmations = confirmations
	self.Receive_At = receive_at
	self.Currency = currency
	TransactionStateMachine.Initial("unconfirm")
	TransactionStateMachine.Trigger("unconfirm", self, db)
}

func (self *PaymentTransaction) Create() {
	db.Create(self)
}

func (self *PaymentTransaction) Save() {
	db.Save(self)
}

// func (self *PaymentTransaction) Get_Deposit() *Deposit {
// 	var deposit Deposit
// 	db.Model(self).Related(&deposit)
// 	return &deposit
// }

func (self *PaymentTransaction) Refresh_Confirmations() (bChanged bool) {
	currency := (&jsons.Currency{}).Find(self.Currency)
	details, err := wallet_service.CoinRPC[currency.Code].(wallet_service.Coind).GetTransaction(self.Txid)
	if err != nil {
		return
	}
	if self.Confirmations != int(details[0].Confirmations) {
		bChanged = true
		self.Confirmations = int(details[0].Confirmations)
	} else {
		bChanged = false
	}
	logger.Info("Confirmations:", self.Confirmations, details[0].Confirmations)
	self.Save()
	return
}

func (self *PaymentTransaction) Min_Confirm() bool {
	// return self.Get_Deposit().Min_Confirm(self.Confirmations)

	// if self.Confirmations >= (&jsons.DepositChannel{}).Find_By_Id(self.Currency).Min_Confirm && self.Confirmations < (&jsons.DepositChannel{}).Find_By_Id(self.Currency).Max_Confirm {
	if self.Confirmations >= 1 && self.Confirmations < 6 {
		return true
	} else {
		return false
	}
}
func (self *PaymentTransaction) Trigger(event string) error {
	return TransactionStateMachine.Trigger(event, self, db)
}

func (self *PaymentTransaction) Notify() {
	data := make(map[string]interface{}, 0)
	data["currency"] = (&jsons.Currency{}).Find(self.Currency).Code
	data["txid"] = self.Txid
	data["address"] = self.Address
	data["amount"] = self.Amount
	data["confirmations"] = self.Confirmations
	AMQPQueue_.Enqueue("deposit", utils.JSONToByteArray(data), nil)
}

func (self *PaymentTransaction) Max_Confirm() bool {
	// if self.Confirmations >= (&jsons.DepositChannel{}).Find_By_Id(self.Currency).Max_Confirm {
	if self.Confirmations >= 6 {
		return true
	} else {
		return false
	}
}

func (self *PaymentTransaction) Get(txid string, txout int) *PaymentTransaction {
	paymentTransaction := &PaymentTransaction{}
	db.Where("txid = ? AND txout = ?", txid, txout).Find(paymentTransaction)
	return paymentTransaction
}

func (self *PaymentTransaction) Get_Not_Confirmed() []*PaymentTransaction {
	paymentTransactions := make([]*PaymentTransaction, 0)
	db.Where("state = ? OR state = ?", "unconfirm", "confirming").Find(&paymentTransactions)
	return paymentTransactions
}

// func (self *PaymentTransaction) Deposit_Accept() {
// 	DepositStateMachine.Trigger("accept", self.Get_Deposit(), db)
// }
