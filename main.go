package main

import (
	"fmt"
	// "math/big"

	"github.com/ADHFMZ7/cryptos/rsa"
	// "github.com/ADHFMZ7/cryptos/aes"
	"github.com/ADHFMZ7/cryptos/util"
)

func main() {

	// var a *big.Int
	// a = util.StringEncode("Ahmad")
	// fmt.Printf("%x\n", a)
	// fmt.Println(util.StringDecode(a))

	pub, priv := rsa.GenerateKeypair(4096)

	// fmt.Println("pub: ", pub)
	// fmt.Println("priv: ", priv)

	message := util.StringEncode("Ahmad Aldasouqi")
	fmt.Println("Ahmad Aldasouqi")

	ciphertext := rsa.Encrypt(message, pub)
	// fmt.Println(ciphertext, "\n\n\n")

	output := rsa.Decrypt(ciphertext, priv)
	// fmt.Println(output, "\n\n\n")

	fmt.Println(util.StringDecode(output))


	// fmt.Println("AES128 Encryption")
	//
	// key := util.HexToBytes("68656C6C6F2D61657300000000000000")
	// in := util.HexToBytes("0011223344550945f554314bbccddeef")
	//
	// util.PrintBytes(in)
	// fmt.Println()
	//
	// ciphertext := aes.Encrypt(in, key)
	//
	// util.PrintBytes(ciphertext)
	// fmt.Println()
	//
	// encrypted := aes.Decrypt(ciphertext, key)
	// util.PrintBytes(encrypted)

}
