package rsa

import (
  "math/big"
  "github.com/ADHFMZ7/cryptos/util"
)

// type Key interface {
//   GetPEM() string
// }

type PublicKey struct {
  Modulus, Exponent *big.Int 
}

type PrivateKey struct {
  Modulus, Exponent *big.Int 
}



func GenerateKeypair(key_bits int) (PublicKey, PrivateKey) {

  q := util.GeneratePrime(key_bits)
  p := util.GeneratePrime(key_bits)

  n := new(big.Int).Mul(q, p) // modulus
  t := totient(p, q)
  e := big.NewInt(65537) // public exponent

  d := big.NewInt(0).ModInverse(e, t) // private exponent

  return PublicKey{n, e}, PrivateKey{n, d}
  
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

// func encrypt(message []byte, pk PublicKey) ([]byte) {
func Encrypt(m *big.Int, pk PublicKey) *big.Int {

  return new(big.Int).Exp(m, pk.Exponent, pk.Modulus)
}

func Decrypt(ciphertext *big.Int, sk PrivateKey) *big.Int {
  return new(big.Int).Exp(ciphertext, sk.Exponent, sk.Modulus)
}
