package jsons

import (
	"github.com/bn_funds/utils"
)

type DepositChannel struct {
	ID          int
	Key         string
	Currency    string
	Sort_Order  int
	Min_Confirm int
	Max_Confirm int
}

func (self *DepositChannel) Init() {

}
func (self *DepositChannel) Find(id int) *DepositChannel {
	channels := self.All()

	for _, item := range channels {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (self *DepositChannel) All() []*DepositChannel {
	byteValue := utils.ReadJSON("deposit_channels")
	channels := make([]*DepositChannel, 0)
	utils.ByteArrayToJSON(byteValue, &channels)
	return channels
}

func (self *DepositChannel) Find_By_Id(id int) *DepositChannel {
	channels := self.All()

	for _, item := range channels {
		if item.ID == id {
			return item
		}
	}

	return nil
}
