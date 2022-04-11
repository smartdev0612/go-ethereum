package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://kovan.infura.io/v3/project_id"
)

func main() {
	// ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// _, err := ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// "be42d3b5f88536bff88e73dc03960a926d06499e"
	// "0ac08f0770859410ce6468386b124061ffc68f40"

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	a1 := common.HexToAddress(("be42d3b5f88536bff88e73dc03960a926d06499e"))
	a2 := common.HexToAddress(("0ac08f0770859410ce6468386b124061ffc68f40"))

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance 1: ", b1)
	fmt.Println("Balance 2: ", b2)
}
