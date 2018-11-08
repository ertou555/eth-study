package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"new-go-eth/token-contract/mytoken"
)

func main() {
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial("/root/1031-dev-eth/ethdev/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	token, err := mytoken.NewMyToken(common.HexToAddress("0xff436696c3995d05a040d6502dfbfc9d99c2d8a2"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	contractName, err := token.Name(nil)
	if err != nil {
		log.Fatalf("query name err:%v", err)
	}
	fmt.Printf("MyToken Name is:%s\n", contractName)
	balance, err := token.BalanceOf(nil, common.HexToAddress("0x8c1b2e9e838e2bf510ec7ff49cc607b718ce8401"))
	if err != nil {
		log.Fatalf("query balance error:%v", err)
	}
	fmt.Printf("0x8c1b2e9e838e2bf510ec7ff49cc607b718ce8401's balance is %s\n", balance)
}
