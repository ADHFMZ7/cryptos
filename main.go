package main

import (
    "fmt"
    // "math/big"

    // "github.com/ADHFMZ7/cryptos/rsa"
    "github.com/ADHFMZ7/cryptos/aes"
    // "github.com/ADHFMZ7/cryptos/util"
    "io/ioutil"
    "log"
    "strings"

)

func main() {
    inputFile := "image.ppm"
    outputFile := "encrypted_image.ppm"
    
    data, err := ioutil.ReadFile(inputFile)
    if err != nil {
        log.Fatalf("Error reading input file: %v", err)
    }

    headerEnd := strings.Index(string(data), "\n255\n") + len("\n255\n")
    header := data[:headerEnd]
    pixelData := data[headerEnd:]

    key := []byte("0123456789abcdef") // Replace with a more secure key

    // Encrypt pixel data using ECB mode
    encryptedData := aes.EncryptCBC(pixelData, key, key)

    // Create new encrypted PPM file with the same header and the encrypted pixel data
    outputData := append(header, encryptedData...)
    err = ioutil.WriteFile(outputFile, outputData, 0644)
    if err != nil {
        log.Fatalf("Error writing output file: %v", err)
    }

    fmt.Printf("Image encrypted successfully. Output written to %s\n", outputFile)
}

	

//
//
// 	fmt.Println("AES128 Encryption")
//
// 	key := util.HexToBytes("68656C6C6F2D61657300000000000000")
// 	in := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()<>?,./;'[]{}:\"")
//
// 	util.PrintBytes(in)
//
// 	IV := []byte("ABCDEFGHIJKLMNOP")
//
// 	ciphertext := aes.EncryptCBC(in, key, IV)
//
// 	fmt.Println()
// 	fmt.Println("\n ciphertext: ")
// 	util.PrintBytes(ciphertext)
// 	fmt.Println()
//
// 	decrypted := aes.DecryptCBC(ciphertext, key, IV)
// 	fmt.Println(string(decrypted))
// }
