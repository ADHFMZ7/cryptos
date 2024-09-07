package main

import (
	"fmt"

	"github.com/ADHFMZ7/cryptos/aes"
)

func main() {

	// key := [16]byte{
	// 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	// 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	// }

	// Byte array for the input
	// in := [16]byte{
	// 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77,
	// 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
	// }
	// aes.PrintState(in)

	// key := [16]byte{
	// 0x00, 0x04, 0x08, 0x0c,  // Column 1
	// 0x01, 0x05, 0x09, 0x0d,  // Column 2
	// 0x02, 0x06, 0x0a, 0x0e,  // Column 3
	// 0x03, 0x07, 0x0b, 0x0f,  // Column 4
	// }

	key := [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// Input in column-major order
	// in := [16]byte{
	// 0x00, 0x44, 0x88, 0xcc,  // Column 1
	// 0x11, 0x55, 0x99, 0xdd,  // Column 2
	// 0x22, 0x66, 0xaa, 0xee,  // Column 3
	// 0x33, 0x77, 0xbb, 0xff,  // Column 4
	// }
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
