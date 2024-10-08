package aes

func EncryptECB(input, key []byte) []byte {

  padding := 16 - (len(input) % 16)
  if padding != 16 {
    for ix := 0; ix < padding; ix++ {
      input = append(input, 0)
    }
  }

  var ciphertext []byte

  for ix := 0; ix < len(input); ix += 16 {
    block := ([16]byte)(input[ix:ix+16])
    for _, cipher_byte := range EncryptBlock(block, ([16]byte)(key)) {
      ciphertext = append(ciphertext, cipher_byte)
    }
  }
   
  return ciphertext
}

func DecryptECB(ciphertext, key []byte) []byte {

  var plaintext []byte

  for ix := 0; ix < len(ciphertext); ix += 16 {
    block := ([16]byte)(ciphertext[ix:ix+16])
    for _, plain_byte := range DecryptBlock(block, ([16]byte)(key)) {
      plaintext = append(plaintext, plain_byte)
    }
  }

  return plaintext
}

func EncryptCBC(input, key, IV []byte) []byte {

  padding := 16 - (len(input) % 16)
  if padding != 16 {
    for ix := 0; ix < padding; ix++ {
      input = append(input, 0)
    }
  }

  var ciphertext []byte

  for ix := 0; ix < len(input); ix += 16 {

    block := ([16]byte)(input[ix:ix+16])
  
    if ix == 0 {
      for iy := range block {
        block[iy] =^ IV[iy]
      }
    } else {
      for iy := range block {
        block[iy] =^ ciphertext[ix - 16 + iy]
      }
    }

    for _, cipher_byte := range EncryptBlock(block, ([16]byte)(key)) {
      ciphertext = append(ciphertext, cipher_byte)
    }
  }
   
  return ciphertext
}

func DecryptCBC(ciphertext, key, IV []byte) []byte {

  var plaintext []byte

  for ix := 0; ix < len(ciphertext); ix += 16 {
    block := ([16]byte)(ciphertext[ix:ix+16])
    for iy, plain_byte := range DecryptBlock(block, ([16]byte)(key)) {
      plaintext = append(plaintext, plain_byte)

      if ix == 0 {
        plaintext[iy] ^= IV[iy]
      } else {
        plaintext[iy] ^= ciphertext[ix - 16 + iy]
      }
    }
  }

  return plaintext

}
