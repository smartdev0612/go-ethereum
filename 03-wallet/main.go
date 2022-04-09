package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(pvk)
	fmt.Println("Private Key: ", hexutil.Encode(pData))

	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println("Public Key: ", hexutil.Encode(puData))

	fmt.Println("Address: ", crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
