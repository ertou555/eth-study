package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"eth-study/go_deploy_mytoken/mytoken"
)

var contract_addr string

func main() {
	contract_addr = "0x2e3ad668820ba4500c78289a6306a51778397c75"
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial("/root/1031-dev-eth/ethdev/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	token, err := mytoken.NewMytoken(common.HexToAddress(contract_addr), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	contractName, err := token.Name(nil)
	if err != nil {
		log.Fatalf("query name err:%v", err)
	}
	fmt.Printf("MyToken Name is:%s\n", contractName)
	balance, err := token.BalanceOf(nil, common.HexToAddress(contract_addr))
	//balance, err := token.BalanceOf(nil, common.HexToAddress("0xa2736ff5581a98fce00ca4ee3917147f1b926f8f"))
	if err != nil {
		log.Fatalf("query balance error:%v", err)
	}
	fmt.Printf("0xa2736ff5581a98fce00ca4ee3917147f1b926f8f's balance is %s\n", balance)
}
