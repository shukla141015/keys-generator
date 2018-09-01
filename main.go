package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
)

var one = big.NewInt(1)

func main() {
	coin := os.Args[1]

	pageNumber := os.Args[2]

	keysPerPage := 128

	switch coin {
	case "btc":
		printBitcoinKeys(pageNumber, keysPerPage)
	case "eth":
		printEthereumKeys(pageNumber, keysPerPage)
	default:
		log.Fatal("Invalid coin type")
	}
}

func printBitcoinKeys(pageNumber string, keysPerPage int) {
	// bitcoinKeys := generateBitcoinKeys(pageNumber, keysPerPage)
}

func printEthereumKeys(pageNumber string, keysPerPage int) {
	ethereumKeys := generateEthereumKeys(pageNumber, keysPerPage)

	length := len(ethereumKeys)

	for i, key := range ethereumKeys {
		fmt.Printf("%v", key)

		if i != length-1 {
			fmt.Print("\n")
		}
	}
}

func makeBigInt(number string) *big.Int {
	i, success := new(big.Int).SetString(number, 10)

	if !success {
		log.Fatal("Failed to create BigInt from string")
	}

	return i
}
