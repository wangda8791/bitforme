package wallet_service

import (
	"log"

	"github.com/bn_funds/jsons"
)

var (
	CoinRPC map[string]interface{}
)

func init_ethrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("eth")
	coind, err := NewEthereumd(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false, currency.Main_address)
	CoinRPC["eth"] = coind

	if err != nil {
		log.Fatalln(err)
	}
}

func init_btcrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("btc")
	coind, err := NewBitcoind(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["btc"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_ltcrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("ltc")
	coind, err := NewLitecoind(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["ltc"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_dogerpc() {
	currency := (&jsons.Currency{}).Find_By_Code("doge")
	coind, err := NewDogecoind(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["doge"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_dashrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("dash")
	coind, err := NewDashd(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["dash"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_bchrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("bch")
	coind, err := NewBitcoinCashd(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["bch"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_qtumrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("qtum")
	coind, err := NewQtumd(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["qtum"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_zecrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("zec")
	coind, err := NewZCashd(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["zec"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_wavesrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("waves")
	coind, err := NewWavesd(currency.Rpc.Host, currency.Rpc.Port, false, "ridethewaves!", currency.Main_address)
	CoinRPC["waves"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_etcrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("etc")
	coind, err := NewEthereumClassicd(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false, currency.Main_address)
	CoinRPC["etc"] = coind

	if err != nil {
		log.Fatalln(err)
	}
}

func init_xlmrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("xlm")
	coind, err := NewStellard(currency.Main_address, currency.Main_seed)
	CoinRPC["xlm"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_xrprpc() {
	currency := (&jsons.Currency{}).Find_By_Code("xrp")
	coind, err := NewXrpd(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["xrp"] = coind
	if err != nil {
		log.Fatalln(err)
	}
	// logger.Info("XRPD:", coind)
	// r, err := coind.GetAccountInfo("rsYeGZRb2d7u2f6LS9XuYTFDzjhm89c3Dp")
	// logger.Info("AccountInfo:", r, err)
}

func init_neorpc() {
	currency := (&jsons.Currency{}).Find_By_Code("neo")
	coind, err := NewNeod(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["neo"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_xmrrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("xmr")
	coind, err := NewMonerod(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.WPort, currency.Rpc.User, currency.Rpc.Password, false, currency.Main_address)
	CoinRPC["xmr"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_adarpc() {
	currency := (&jsons.Currency{}).Find_By_Code("ada")
	coind, err := NewCardanod(currency.Rpc.Host, currency.Rpc.Port, true, currency.Main_address, currency.Spending_pass, currency.Account_index)
	CoinRPC["ada"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_trxrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("trx")
	coind, err := NewTrond(currency.Rpc.Host, currency.Rpc.Port, false, currency.Main_address, currency.Priv_key)
	CoinRPC["trx"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func init_vtcrpc() {
	currency := (&jsons.Currency{}).Find_By_Code("vtc")
	coind, err := NewVertcoind(currency.Rpc.Host, currency.Rpc.Port, currency.Rpc.User, currency.Rpc.Password, false)
	CoinRPC["vtc"] = coind
	if err != nil {
		log.Fatalln(err)
	}
}

func Init_Rpc() {
	CoinRPC = make(map[string]interface{}, 0)
	// btc styled rpc
	init_btcrpc()
	init_ltcrpc()
	init_dogerpc()
	init_dashrpc()
	init_bchrpc()
	init_qtumrpc()
	init_zecrpc()
	init_vtcrpc()

	//ethereum styled rpc
	init_ethrpc()
	init_wavesrpc()
	init_etcrpc()
	init_xlmrpc()
	init_neorpc()
	init_xmrrpc()
	init_adarpc()
	init_trxrpc()
}
