package main

import (
  "fmt"
  "github.com/ADHFMZ7/cryptos/aes"
)

func main() {

  state := [16]byte {0x74, 0xc5, 0xdf, 0x3c, 
                      0x6c, 0x1e,  0x93, 0x62,
                      0xe1, 0xdd, 0x79, 0xb0,
                      0x09, 0x3b, 0xc7, 0xe7}

  // state := [16]byte {0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

  state = aes.SubBytes(state)
  // state = aes.ShiftRows(state)

  // fmt.Println(state[:4])
  // fmt.Println(state[4:8])
  // fmt.Println(state[8:12])
  // fmt.Println(state[12:])

  for _, b := range state[:4] {
    fmt.Printf("%02x ", b)
  }
  fmt.Println()

  for _, b := range state[4:8] {
    fmt.Printf("%02x ", b)
  }
  fmt.Println()

  for _, b := range state[8:12] {
    fmt.Printf("%02x ", b)
  }
  fmt.Println()

  for _, b := range state[12:] {
    fmt.Printf("%02x ", b)
  }
  fmt.Println()

}
