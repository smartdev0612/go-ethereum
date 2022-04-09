package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/4badaf12963a4dd28d67d5a57e71a75f"
var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block:%v", err)
	}
	fmt.Println("The block number: ", block.Number())

	addr := "0x439accf33b62db96E295e9E184CD55b8D3B9e9EC"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance:%v", err)
	}
	fmt.Println("The balance:", balance)

	// 1 ether = 10 ^ 18 wei
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println(fBalance)

	// 10*10*10*... 18
	balanceEther := new(big.Float).Quo(fBalance, big.NewFloat((math.Pow10(18))))
	fmt.Println(balanceEther)
}
