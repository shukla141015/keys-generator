package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	coin := os.Args[1]

	keysPerPage := 128

	switch coin {
	case "btc":
		printBitcoinKeys(os.Args[2], keysPerPage)
	case "btc-search":
		printBtcWifSearch(os.Args[2], keysPerPage)
	case "eth":
		printEthereumKeys(os.Args[2], keysPerPage)
	case "eth-search":
		printEthPrivateKeySearch(os.Args[2], keysPerPage)
	default:
		log.Fatal("Invalid coin type")
	}
}

func printBitcoinKeys(pageNumber string, keysPerPage int) {
	bitcoinKeys := generateBitcoinKeys(pageNumber, keysPerPage)

	length := len(bitcoinKeys)

	for i, key := range bitcoinKeys {
		fmt.Printf("%v", key)

		if i != length-1 {
			fmt.Print("\n")
		}
	}
}

func printBtcWifSearch(wif string, keysPerPage int) {
	pageNumber := findBtcWifPage(wif, keysPerPage)

	fmt.Printf("%v", pageNumber)
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

func printEthPrivateKeySearch(privateKey string, keysPerPage int) {
	pageNumber := findEthPrivateKeyPage(privateKey, keysPerPage)

	fmt.Printf("%v", pageNumber)
}
