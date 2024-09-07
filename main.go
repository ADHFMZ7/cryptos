package main

import (
	"fmt"

	"github.com/ADHFMZ7/cryptos/aes"
)

func main() {

	// state := [16]byte { 0xd4, 0xc5, 0xdf, 0x3c,
	//                     0xbf, 0x1e,  0x93, 0x62,
	//                     0x5d, 0xdd, 0x79, 0xb0,
	//                     0x30, 0x3b, 0xc7, 0xe7}
	//

	// state := [16]byte { 0x74, 0x6c, 0xe1, 0x09,
	//                     0xc5, 0x1e, 0xdd, 0x6b,
	//                     0xdf, 0x93, 0x79, 0xc7,
	//                     0x3c, 0x62, 0xb0, 0xe7}

	// state := [16]byte {0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

	state := [16]byte{0xd4, 0xbf, 0x5d, 0x30,
		0xc5, 0x1e, 0xdd, 0x6b,
		0xdf, 0x93, 0x79, 0xc7,
		0x3c, 0x62, 0xb0, 0xe7}

	aes.PrintState(state)
	fmt.Println("\n==========")

	state = aes.MixColumns(state)
	// state = aes.ShiftRows(state)
	// state = aes.SubBytes(state)


	aes.PrintState(state)

}
