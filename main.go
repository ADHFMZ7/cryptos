package main

import (
	"fmt"

	"github.com/ADHFMZ7/cryptos/aes"
)

func main() {

	key := [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	in := [16]byte{
		0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77,
		0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
	}

	schedule := aes.KeySchedule(key)

	for ix := range schedule {
		fmt.Println("\n ", ix)
		aes.PrintState(schedule[ix])
		fmt.Println("\n ")
	}

	ciphertext := aes.Encrypt(in, schedule)

	aes.PrintState(ciphertext)
}
