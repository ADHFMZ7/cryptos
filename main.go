package main

import (
	"fmt"
	// "math/big"

	// "github.com/ADHFMZ7/cryptos/rsa"
	"github.com/ADHFMZ7/cryptos/aes"
	"github.com/ADHFMZ7/cryptos/util"
)

func main() {


	fmt.Println("AES128 Encryption")

	key := util.HexToBytes("68656C6C6F2D61657300000000000000")
	in := util.HexToBytes("0011223344550945f554314bbccddeefaa")

	util.PrintBytes(in)

	ciphertext := aes.EncryptECB(in, key)

	fmt.Println()
	fmt.Println("\n ciphertext: ")
	util.PrintBytes(ciphertext)
	fmt.Println()

	decrypted := aes.DecryptECB(ciphertext, key)
	util.PrintBytes(decrypted)
}
