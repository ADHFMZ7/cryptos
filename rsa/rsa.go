package rsa

import (
  "math/big"
  "crypto/rand"
)

type PublicKey struct {

  

}

func GeneratePrime(bits int) *big.Int {

  num, err := rand.Prime(rand.Reader, bits)

  if err != nil {
    panic(err)
  }

  return num
}


func GenerateKeypair(key_bits int) (*big.Int, *big.Int) {

  q := GeneratePrime(key_bits)
  p := GeneratePrime(key_bits)

  n := new(big.Int).Add(q, p)
  t := totient(p, q)
  e := 65537

  

}

func totient(p, q *big.Int) *big.Int {

  return lcm(
    new(big.Int).Sub(p, big.NewInt(1)), 
    new(big.Int).Sub(q, big.NewInt(1)),
  ) 
}

func lcm(a, b *big.Int) *big.Int {
 
  numerator := big.NewInt(0)
  numerator.Mul(a, b)

  return numerator.Div(numerator, new(big.Int).GCD(nil, nil, a, b))
}

func encrypt(message []byte, pk big.Int) ([]byte) {

   
  


  return make([]byte, 0) 
}

func decrypt(ciphertext []byte, sk big.Int) ([]byte) {
  return make([]byte, 0)
}
