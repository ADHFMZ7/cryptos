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
	in := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()<>?,./;'[]{}:\"")

	util.PrintBytes(in)

	IV := []byte("ABCDEFGHIJKLMNOP")

	ciphertext := aes.EncryptCBC(in, key, IV)

	fmt.Println()
	fmt.Println("\n ciphertext: ")
	util.PrintBytes(ciphertext)
	fmt.Println()

	decrypted := aes.DecryptCBC(ciphertext, key, IV)
	fmt.Println(string(decrypted))
}
