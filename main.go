package main

import (
	"fmt"

	"github.com/ADHFMZ7/cryptos/aes"
	"github.com/ADHFMZ7/cryptos/util"
)

func main() {

	fmt.Println("AES128 Encryption")

	key := util.HexToBytes("0001020a4935060708090a0b0c0d0e0f")
	in := util.HexToBytes("0011223344550945f554314bbccddeef")

	util.PrintBytes(in)
	fmt.Println()

	ciphertext := aes.Encrypt(in, key)

	util.PrintBytes(ciphertext)
	fmt.Println()

	encrypted := aes.Decrypt(ciphertext, key)
	util.PrintBytes(encrypted)

}
