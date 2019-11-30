package worker

import (
	"context"
	"fmt"

	"github.com/bn_funds/jsons"
	"github.com/bn_funds/models"
	"github.com/bn_funds/services/wallet_service"
	"github.com/bn_funds/utils"
	"github.com/google/logger"
	"github.com/stellar/go/clients/horizon"
)

type BlockCheck struct {
}

func (self *BlockCheck) Init() {
	currencies := [...]string{"eth", "etc", "waves", "neo", "xmr", "ada", "trx"}
	for _, currency := range currencies {
		currency_id := (&jsons.Currency{}).Find_By_Code(currency).ID
		blockCheck := (&models.BlockCheck{}).GetCheckedLastBlock(currency)
		blockNumber, err := wallet_service.CoinRPC[currency].(wallet_service.Coind).GetBlockNumber()
		if err != nil {
			blockNumber = 0
		}
		if blockCheck.ID == 0 {
			blockCheck.Init(currency_id, int(blockNumber))
			blockCheck.Create()
		}

		if utils.GetEnv("PREV_BLOCK_CHECK", "true") == "false" {
			blockCheck.Last = int(blockNumber)
			blockCheck.Save()
		}
	}

}

func (self *BlockCheck) ProcessMTM() {
	currencies := [...]string{"eth", "etc", "waves", "trx"}
	for _, currency := range currencies {
		blockCheck := (&models.BlockCheck{}).GetCheckedLastBlock(currency)
		logger.Info(currency, blockCheck)
		txids, latestBlockNumber, err := wallet_service.CoinRPC[currency].(wallet_service.Coind).GetDepositTransactions(blockCheck.Last)
		if err != nil {
			continue
		}
		for _, txid := range txids {
			(&DepositCoin{}).Process2(currency, txid)
		}
		blockCheck.Last = int(latestBlockNumber)
		blockCheck.Save()
	}
}

func (self *BlockCheck) ProcessMTO() {
	currencies := [...]string{"neo", "xmr", "ada"}

	for _, currency := range currencies {
		blockCheck := (&models.BlockCheck{}).GetCheckedLastBlock(currency)
		logger.Info(currency, blockCheck)
		txids, latestBlockNumber, err := wallet_service.CoinRPC[currency].(wallet_service.Coind).GetDepositTransactions(blockCheck.Last)
		if err != nil {
			continue
		}
		for _, txid := range txids {
			(&DepositCoin{}).Process(currency, txid)
		}
		blockCheck.Last = int(latestBlockNumber)
		blockCheck.Save()
	}
}

func (self *BlockCheck) ProcessXLM() {
	address := wallet_service.CoinRPC["xlm"].(wallet_service.Coind).GetMainAddress()
	ctx := context.Background()

	cursor := horizon.Cursor("now")

	fmt.Println("Waiting for a payment...")

	err := horizon.DefaultTestNetClient.StreamPayments(ctx, address, &cursor, func(payment horizon.Payment) {
		fmt.Println("Payment:", payment.TransactionHash)
		(&DepositCoin{}).Process("xlm", payment.TransactionHash)

	})

	if err != nil {
		panic(err)
	}

}
