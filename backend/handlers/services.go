package handlers

import (
	"fmt"
	"supply-chain/utils/SupplyChain"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
)

// fisco 客户端 API 调用
var fiscoCli *client.Client

// 编译后的智能合约 API，这个才是调用自己的智能合约接口的 API
var contractAPI *SupplyChain.SupplyChainSession

// 初始化 fisco client 和 编译好的智能合约的 API
func SetFsicoCli(cli *client.Client) {
	fiscoCli = cli
	addrString := "0xaaa7a49d4d49fbbb08e60d86a6890499b25757ff"
	fmt.Println("contract addr: ", addrString)
	contractAddr := common.HexToAddress(addrString)
	instance, _ := SupplyChain.NewSupplyChain(contractAddr, fiscoCli)
	contractAPI = &SupplyChain.SupplyChainSession{
		Contract:     instance,
		CallOpts:     *fiscoCli.GetCallOpts(),
		TransactOpts: *fiscoCli.GetTransactOpts(),
	}
}
