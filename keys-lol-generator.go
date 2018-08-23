package main

import (
	"fmt"
    "os"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

const KeysPerPage = 128

var (
	// Total amount of seeds
	total = new(big.Int).SetBytes([]byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE,
		0xBA, 0xAE, 0xDC, 0xE6, 0xAF, 0x48, 0xA0, 0x3B, 0xBF, 0xD2, 0x5E, 0x8C, 0xD0, 0x36, 0x41, 0x40,
	})
	
	one = big.NewInt(1)
)

type Key struct {
	private      string
	number       string
	compressed   string
	uncompressed string
}

func generateKeys(count *big.Int) (keys [KeysPerPage]Key, length int) {
	var padded [32]byte

	var i int

	for i = 0; i < KeysPerPage; i++ {
		
		count.Add(count, one)

		// Check to make sure we're not out of range
		if count.Cmp(total) > 0 {
			break
		}

		// Copy count value's bytes to padded slice
		copy(padded[32-len(count.Bytes()):], count.Bytes())

		// Get private and public keys
		privKey, public := btcec.PrivKeyFromBytes(btcec.S256(), padded[:])

		// Get compressed and uncompressed addresses for public key
		caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)

		// Encode addresses
		wif, _ := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, false)
		keys[i].private = wif.String()
		keys[i].number = count.String()
		keys[i].compressed = caddr.EncodeAddress()
		keys[i].uncompressed = uaddr.EncodeAddress()
	}

	return keys, i
}

func main() {
    // Get the first argument
	pageNumber := os.Args[1]
    
    page, success := new(big.Int).SetString(pageNumber, 0)
	
    if !success {		
		return
	}
    
    previousPage := new(big.Int).Sub(page, one)

	// Calculate our starting seed based on the current page number
	firstSeed := new(big.Int).Mul(previousPage, big.NewInt(KeysPerPage))
	
	keys, length := generateKeys(firstSeed)

	for i := 0; i < length; i++ {
		key := keys[i]
		
        fmt.Print(key.private + " " + key.number + " " +  key.uncompressed + " " + key.compressed + "\n")
	}
}
