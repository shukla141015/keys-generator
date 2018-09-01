package main

import (
	"log"
	"math/big"
)

var one = big.NewInt(1)

func makeBigInt(number string) *big.Int {
	i, success := new(big.Int).SetString(number, 10)

	if !success {
		log.Fatal("Failed to create BigInt from string")
	}

	return i
}
