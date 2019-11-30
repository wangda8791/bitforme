package models

import (
	"github.com/google/logger"
	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
)

func Init_State_Machine() {
	init_transation_state_machine()
	// init_deposit_state_machine()
}

func init_transation_state_machine() {
	TransactionStateMachine = transition.New(&PaymentTransaction{})
	TransactionStateMachine.State("unconfirm")
	TransactionStateMachine.State("confirming")
	TransactionStateMachine.State("confirmed")

	TransactionStateMachine.Event("check_confirming").To("confirming").From("unconfirm").Before(func(transaction interface{}, tx *gorm.DB) error {
		return nil
	}).After(func(transaction interface{}, tx *gorm.DB) error {
		transaction.(*PaymentTransaction).Save()
		logger.Info("check_confirming:", transaction.(*PaymentTransaction).Confirmations)
		// transaction.(*PaymentTransaction).Notify()
		// logger.Info("after confirming")
		return nil
	})

	TransactionStateMachine.Event("check_confirmed").To("confirmed").From("unconfirm", "confirming").Before(func(transaction interface{}, tx *gorm.DB) error {
		return nil
	}).After(func(transaction interface{}, tx *gorm.DB) error {
		transaction.(*PaymentTransaction).Save()
		logger.Info("check_confirmed:", transaction.(*PaymentTransaction).Confirmations)
		transaction.(*PaymentTransaction).Notify()
		return nil
	})
}

// func init_deposit_state_machine() {
// 	DepositStateMachine = transition.New(&Deposit{})
// 	DepositStateMachine.State("submitting")
// 	DepositStateMachine.State("cancelled")
// 	DepositStateMachine.State("submitted")
// 	DepositStateMachine.State("rejected")
// 	DepositStateMachine.State("accepted")
// 	DepositStateMachine.State("checked")
// 	DepositStateMachine.State("warning")

// 	DepositStateMachine.State("submitting").Enter(func(deposit interface{}, tx *gorm.DB) error {
// 		deposit.(*Deposit).Set_Fee()
// 		return nil
// 	}).Exit(func(deposit interface{}, tx *gorm.DB) error {

// 		return nil
// 	})

// 	DepositStateMachine.Event("submit").To("submitted").From("submitting").After(func(deposit interface{}, tx *gorm.DB) error {
// 		deposit.(*Deposit).Save()
// 		return nil
// 	})

// 	DepositStateMachine.Event("cancel").To("cancelled").From("submitting").After(func(deposit interface{}, tx *gorm.DB) error {
// 		deposit.(*Deposit).Save()
// 		return nil
// 	})

// 	DepositStateMachine.Event("reject").To("rejected").From("submitted").After(func(deposit interface{}, tx *gorm.DB) error {
// 		deposit.(*Deposit).Save()
// 		return nil
// 	})

// 	DepositStateMachine.Event("accept").To("accepted").From("submitted").After(func(deposit interface{}, tx *gorm.DB) error {
// 		deposit.(*Deposit).Save()
// 		deposit.(*Deposit).Do()
// 		deposit.(*Deposit).Send_Mail()
// 		deposit.(*Deposit).Send_SMS()
// 		return nil
// 	})

// 	DepositStateMachine.Event("check").To("checked").From("accepted").After(func(deposit interface{}, tx *gorm.DB) error {
// 		deposit.(*Deposit).Save()
// 		return nil
// 	})

// 	DepositStateMachine.Event("warn").To("warning").From("accepted").After(func(deposit interface{}, tx *gorm.DB) error {
// 		deposit.(*Deposit).Save()
// 		return nil
// 	})
// }
