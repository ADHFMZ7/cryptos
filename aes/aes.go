package aes

// S-box lookup table
func subBytes(state [16]byte) [16]byte {

	for ix, value := range state {
		state[ix] = sbox0[value]
	}
	return state
}

// Inverse of subBytes
func subBytesInv(state [16]byte) [16]byte {
	for ix, value := range state {
		state[ix] = sboxinv[value]
	}
	return state

}

// Performs left circular shift on each row of the state matrix
func shiftRows(state [16]byte) [16]byte {

	var temp [3]byte

	temp[0] = state[1]
	state[1] = state[5]
	state[5] = state[9]
	state[9] = state[13]
	state[13] = temp[0]

	temp[0] = state[2]
	temp[1] = state[6]
	state[2] = state[10]
	state[6] = state[14]
	state[10] = temp[0]
	state[14] = temp[1]

	temp[0] = state[3]
	temp[1] = state[7]
	temp[2] = state[11]
	state[3] = state[15]
	state[7] = temp[0]
	state[11] = temp[1]
	state[15] = temp[2]

	return state
}

// Inverse of shiftRows
func shiftRowsInv(state [16]byte) [16]byte {

	var temp [3]byte

	temp[0] = state[13]
	state[13] = state[9]
	state[9] = state[5]
	state[5] = state[1]
	state[1] = temp[0]

	temp[0] = state[14]
	temp[1] = state[10]
	state[14] = state[6]
	state[10] = state[2]
	state[6] = temp[0]
	state[2] = temp[1]

	temp[0] = state[15]
	temp[1] = state[11]
	temp[2] = state[7]
	state[15] = state[3]
	state[11] = temp[0]
	state[7] = temp[1]
	state[3] = temp[2]

	return state
}

func mixColumn(column [4]*byte) {
	temp := [4]byte{0, 0, 0, 0}

	temp[0] = GFMul(0x02, *column[0]) ^ GFMul(0x03, *column[1]) ^ *column[2] ^ *column[3]
	temp[1] = *column[0] ^ GFMul(0x02, *column[1]) ^ GFMul(0x03, *column[2]) ^ *column[3]
	temp[2] = *column[0] ^ *column[1] ^ GFMul(0x02, *column[2]) ^ GFMul(0x03, *column[3])
	temp[3] = GFMul(0x03, *column[0]) ^ *column[1] ^ *column[2] ^ GFMul(0x02, *column[3])

	*column[0] = temp[0]
	*column[1] = temp[1]
	*column[2] = temp[2]
	*column[3] = temp[3]
}

func mixColumns(state [16]byte) [16]byte {

	for ix := 0; ix < 4; ix++ {
		mixColumn([4]*byte{&state[ix*4], &state[ix*4+1], &state[ix*4+2], &state[ix*4+3]})
	}

	return state
}

// Inverse of mixColumns
func mixColumnsInv(state [16]byte) [16]byte {

	for ix := 0; ix < 4; ix++ {
		mixColumnInv([4]*byte{&state[ix*4], &state[ix*4+1], &state[ix*4+2], &state[ix*4+3]})
	}

	return state
}

// Inverse of mixColumn
func mixColumnInv(column [4]*byte) {

	temp := [4]byte{0, 0, 0, 0}

	temp[0] = GFMul(0x0E, *column[0]) ^ GFMul(0x0B, *column[1]) ^ GFMul(0x0D, *column[2]) ^ GFMul(0x09, *column[3])
	temp[1] = GFMul(0x09, *column[0]) ^ GFMul(0x0E, *column[1]) ^ GFMul(0x0B, *column[2]) ^ GFMul(0x0D, *column[3])
	temp[2] = GFMul(0x0D, *column[0]) ^ GFMul(0x09, *column[1]) ^ GFMul(0x0E, *column[2]) ^ GFMul(0x0B, *column[3])
	temp[3] = GFMul(0x0B, *column[0]) ^ GFMul(0x0D, *column[1]) ^ GFMul(0x09, *column[2]) ^ GFMul(0x0E, *column[3])

	*column[0] = temp[0]
	*column[1] = temp[1]
	*column[2] = temp[2]
	*column[3] = temp[3]
}

func rotWord(column [4]byte) [4]byte {

	var temp byte

	temp = column[0]
	column[0] = column[1]
	column[1] = column[2]
	column[2] = column[3]
	column[3] = temp

	return column
}

// Substitutes each byte in a 4-byte column using the S-box lookup table
func subWord(column [4]byte) [4]byte {

	for ix, val := range column {
		column[ix] = sbox0[val]
	}

	return column
}

// Generates a round constant for the key schedule
func rconWord(column [4]byte, round int) [4]byte {

	// if round > 11 {
	// }

	column[0] ^= rcon[round-1]
	return column
}

// Adds a 16-byte round key to a 16-byte state
func addKey(state, key [16]byte) [16]byte {

	for ix, val := range key {
		state[ix] ^= val
	}
	return state
}

// Expands a 16-byte secret key into 11 16-byte round keys
func KeySchedule(secretKey [16]byte) [11][16]byte {

	var keys_out [11][16]byte

	keys_out[0] = secretKey

	for round := 1; round < 11; round++ {

		secretKey = keys_out[round-1]

		transform := [4]byte{secretKey[12], secretKey[13], secretKey[14], secretKey[15]}
		transform = rotWord(transform)
		transform = subWord(transform)
		transform = rconWord(transform, round)

		wordAdd([4]*byte{&secretKey[0], &secretKey[1], &secretKey[2], &secretKey[3]}, transform)

		wordAdd([4]*byte{&secretKey[4], &secretKey[5], &secretKey[6], &secretKey[7]},
			[4]byte{secretKey[0], secretKey[1], secretKey[2], secretKey[3]})

		wordAdd([4]*byte{&secretKey[8], &secretKey[9], &secretKey[10], &secretKey[11]},
			[4]byte{secretKey[4], secretKey[5], secretKey[6], secretKey[7]})

		wordAdd([4]*byte{&secretKey[12], &secretKey[13], &secretKey[14], &secretKey[15]},
			[4]byte{secretKey[8], secretKey[9], secretKey[10], secretKey[11]})

		keys_out[round] = secretKey
	}
	return keys_out
}

// Encrypts a 16-byte block using a 16-byte private key
func EncryptBlock(state [16]byte, privateKey [16]byte) [16]byte {

	keySchedule := KeySchedule(privateKey)

	state = addKey(state, keySchedule[0])

	for ix := 1; ix < 10; ix++ {

		state = subBytes(state)
		state = shiftRows(state)
		state = mixColumns(state)
		state = addKey(state, keySchedule[ix])

	}

	state = subBytes(state)
	state = shiftRows(state)
	state = addKey(state, keySchedule[10])

	return state
}

// Decrypts a 16-byte ciphertext using a 16-byte private key
func DecryptBlock(ciphertext [16]byte, privateKey [16]byte) [16]byte {

	keySchedule := KeySchedule(privateKey)

	ciphertext = addKey(ciphertext, keySchedule[10])
	ciphertext = shiftRowsInv(ciphertext)
	ciphertext = subBytesInv(ciphertext)

	for ix := 9; ix > 0; ix-- {
		ciphertext = addKey(ciphertext, keySchedule[ix])
		ciphertext = mixColumnsInv(ciphertext)
		ciphertext = shiftRowsInv(ciphertext)
		ciphertext = subBytesInv(ciphertext)
	}

	ciphertext = addKey(ciphertext, keySchedule[0])
	return ciphertext
}

// Addition of two words in GF(2^8)
func wordAdd(a [4]*byte, b [4]byte) {
	for ix, val := range b {
		*a[ix] ^= val
	}
}

// Multiplication in GF(2^8)
func GFMul(a, b byte) byte {
	var result byte

	for ix := 0; ix < 8; ix++ {

		if (b & 1) != 0 {
			result ^= a
		}

		if a&0x80 != 0 {
			a = (a << 1) ^ 0x1b
		} else {
			a <<= 1
		}

		b >>= 1
	}
	return result
}


