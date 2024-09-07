package util

import "fmt"

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
