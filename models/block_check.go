package models

import (
	"github.com/bn_funds/jsons"
	"github.com/jinzhu/gorm"
)

type BlockCheck struct {
	gorm.Model

	Currency int
	Last     int
}

func (self *BlockCheck) Init(currency, last int) {

	self.Currency = currency
	self.Last = last

}

func (self *BlockCheck) Create() {
	db.Create(self)
}

func (self *BlockCheck) Save() {
	db.Save(self)
}

func (self *BlockCheck) GetCheckedLastBlock(currency string) *BlockCheck {
	currency_id := (&jsons.Currency{}).Find_By_Code(currency).ID

	blockCheck := &BlockCheck{}
	db.Where("currency = ?", currency_id).Find(blockCheck)

	return blockCheck
}
