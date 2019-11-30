package worker

import (
	"github.com/bn_funds/models"
)

type PaymentTransaction struct {
}

func (self *PaymentTransaction) Init() {

}

func (self *PaymentTransaction) Process() {
	paymentTransactions := (&models.PaymentTransaction{}).Get_Not_Confirmed()
	for _, v := range paymentTransactions {

		bChanged := v.Refresh_Confirmations()
		if bChanged {
			if v.Min_Confirm() {
				v.Trigger("check_confirming")
				v.Notify()
			} else if v.Max_Confirm() {
				v.Trigger("check_confirmed")
			}
		}
	}
}
