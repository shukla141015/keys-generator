package main

import (
    "fmt"
    "math/big"
    "os"
)

const keysPerPage = 128

func main() {
    pageNumber := os.Args[1]

    basePage := new(big.Int).Sub(makeBigInt(pageNumber), one)

    firstSeed := new(big.Int).Add(new(big.Int).Mul(basePage, big.NewInt(keysPerPage)), one)

    keys := generateBitcoinKeys(firstSeed, keysPerPage)

    for _, key := range keys {
        fmt.Print(key.private + " " + key.uncompressed + " " + key.compressed + "\n")
    }
}

func makeBigInt(number string) *big.Int {
    i, success := new(big.Int).SetString(number, 10)

    if !success {
        panic("Failed to create BigInt from string")
    }

    return i
}
