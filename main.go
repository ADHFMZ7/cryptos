package main

import (
	"fmt"
	"math/big"

	"github.com/ADHFMZ7/cryptos/rsa"
	// "github.com/ADHFMZ7/cryptos/aes"
	// "github.com/ADHFMZ7/cryptos/util"
)

func main() {

	pub, priv := rsa.GenerateKeypair(1024)

	fmt.Println("pub: ", pub)
	fmt.Println("priv: ", priv)

	message := big.NewInt(0)

	message.SetString("33791961580247581621785430719135749904176771115131644124782298049574482190116219659032821583789601179936694916510480822676451949507956785453889578492607999761524741586373725040287757073033386366446512971473881364063101874372135095694432981368154995093751716608693163816670448964807012819160343860257510576144", 10)
	fmt.Println(message, "\n\n\n")

	ciphertext := rsa.Encrypt(message, pub)
	fmt.Println(ciphertext, "\n\n\n")

	output := rsa.Decrypt(ciphertext, priv)
	fmt.Println(output, "\n\n\n")



	//
	// fmt.Println("AES128 Encryption")
	//
	// key := util.HexToBytes("0001020a4935060708090a0b0c0d0e0f")
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
	//
}
