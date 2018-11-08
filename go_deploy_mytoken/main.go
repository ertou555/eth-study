package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"time"
	// GOPATH can find
	"eth-study/go_deploy_mytoken/mytoken"
)

const key = `
{"address":"a2736ff5581a98fce00ca4ee3917147f1b926f8f","crypto":{"cipher":"aes-128-ctr","ciphertext":"13ab8393f9f4af9229b823a46c4917841954ee33f6b07dfa1f21b5b4e0d6e119","cipherparams":{"iv":"eb40e9d4650da24c9d38b576b4ab5b48"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"67de2bba712be196a2d33572a29fd097e3e9672b23f47fa36c1741a809b9bbf0"},"mac":"6e4556b2dcebdc056206bc7ab817a0b9f97164ae39c5b09ee23b11791481c919"},"id":"b6b78d2f-916c-44a4-bb8d-785f9fbcad45","version":3}
`

func main() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("http://localhost:8545/")
	//	conn, err := ethclient.Dial("/root/1031-dev-eth/ethdev/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := mytoken.DeployMytoken(auth, conn, big.NewInt(9651), "Contracts in Go!!!", 0, "Go!")
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	startTime := time.Now()
	fmt.Printf("TX start @:%s", time.Now())

	ctx := context.Background()
	addressAfterMined, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}
	fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
	if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
		log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
	}
	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
}
