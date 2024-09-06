package main

import (
  "fmt"
  "github.com/ADHFMZ7/cryptos/aes"
)

func main() {

  fmt.Println("Ahmad")

  state := [16]byte {0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

  state = aes.ShiftRows(state)

  fmt.Println(state[:4])
  fmt.Println(state[4:8])
  fmt.Println(state[8:12])
  fmt.Println(state[12:])

}
