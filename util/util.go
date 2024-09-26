package util

import (
	"fmt"
	"crypto/rand"
	"math/big"
	// "strings"
)

func StringEncode(input string) *big.Int {
	// for now this only works with ASCII encoded strings

	ret := big.NewInt(0)

	for _, char := range input {
		ret.Lsh(ret, 8)
		ret.Add(ret, big.NewInt(int64(char)))
	}
	return ret
}

func StringDecode(encoded_int *big.Int) string {
	// Estimate the number of bytes needed to store the encoded_int
	byteLen := (encoded_int.BitLen() + 7) / 8
	result := make([]byte, byteLen)
	
	mod := big.NewInt(0)
	for i := byteLen - 1; i >= 0; i-- {
		encoded_int.DivMod(encoded_int, big.NewInt(0xff + 1), mod)
		result[i] = mod.Bytes()[0]
	}
	return string(result)
}

func GeneratePrime(bits int) *big.Int {

	num, err := rand.Prime(rand.Reader, bits)

	if err != nil {
		panic(err)
	}

	return num
}

func HexToBytes(hex string) [16]byte {
	var bytes [16]byte
	for i := 0; i < 16; i++ {
		fmt.Sscanf(hex[2*i:2*i+2], "%x", &bytes[i])
	}
	return bytes
}

// Utility function to print state matrix in row-order
func PrintBytes(state [16]byte) {
	for _, val := range state {
		fmt.Printf("%02x", val)
	}
	fmt.Println()
}

// Utility function to print state matrix in column-order
func PrintState(state [16]byte) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			index := i + 4*j
			fmt.Printf("%02x ", state[index])
		}
		fmt.Println()
	}
	fmt.Println()
}
