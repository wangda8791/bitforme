package jsons

import (
	"github.com/bn_funds/utils"
)

type Currency struct {
	ID                 int
	Key                string
	Code               string
	Name               string
	Symbol             string
	Coin               bool
	Precision          int
	Quick_withdraw_max string
	Rpc                struct {
		Host     string
		Port     int
		WPort    int
		User     string
		Password string
	}
	Main_address  string
	Priv_key      string
	Main_seed     string
	Account_index string
	Spending_pass string
	Blockchain    string
	Address_url   string
	Assets        Asset
}

type Asset struct {
	Accounts []Bank
}

type Bank struct {
	Bank     string
	Address  string
	Password string
	Tel      string
}

func (self *Currency) Init() {

}
func (self *Currency) Find(id int) *Currency {
	currencies := self.All()

	for _, item := range currencies {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (self *Currency) All() []*Currency {
	byteValue := utils.ReadJSON("currencies")
	currencies := make([]*Currency, 0)
	utils.ByteArrayToJSON(byteValue, &currencies)
	return currencies
}

func (self *Currency) Find_By_Code(code string) *Currency {
	currencies := self.All()

	for _, item := range currencies {
		if item.Code == code {
			return item
		}
	}

	return nil
}

func (self *Currency) Get_Fiats() []*Currency {
	c_all := self.All()

	v := CVector{}

	for _, item := range c_all {
		if item.Coin == false {
			v.Add(item)
		}
	}

	return v.vector
}

func (self *Currency) Get_Coins() []*Currency {
	c_all := self.All()

	v := CVector{}

	for _, item := range c_all {
		if item.Coin == true {
			v.Add(item)
		}
	}

	return v.vector
}

type CVector struct {
	vector []*Currency
}

func (self *CVector) Init(item []*Currency) {
	self.vector = make([]*Currency, 0)
}

// func (self *Vector) Empty() bool {
// 	if len(self.vector) == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func (self *CVector) Add(item *Currency) {
	self.vector = append(self.vector, item)
}
