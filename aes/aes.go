package aes

import "fmt"

func generate_round_keys(secret_key int) {

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
}

// Transformations

func SubBytes(state [16]byte) [16]byte {

	for ix, value := range state {
		state[ix] = sbox0[value]
	}
	return state
}

func ShiftRows(state [16]byte) [16]byte {

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

func MixColumn(column [4]*byte) {
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

func MixColumns(state [16]byte) [16]byte {

	for ix := 0; ix < 4; ix++ {
		MixColumn([4]*byte{&state[ix*4], &state[ix*4+1], &state[ix*4+2], &state[ix*4+3]})
	}

	return state
}

func RotWord(column [4]byte) [4]byte {

	var temp byte

	temp = column[0]
	column[0] = column[1]
	column[1] = column[2]
	column[2] = column[3]
	column[3] = temp

	return column
}

func addKey() {

}

func AesKeySchedule(secretKey [16]byte) {

	var keys_out [11][16]byte

	keys_out[0] = secretKey

	for ix := 1; ix <= 11; ix++ {

	}

}

func aesEncrypt(state, keySchedule [16]byte) {

	for ix := 0; ix < 9; ix++ {

		state = SubBytes(state)
		state = ShiftRows(state)
		state = MixColumns(state)
		// state = AddKey(state, key)

	}

	state = SubBytes(state)
	state = ShiftRows(state)
	// state = AddKey(state, key)

}

func GFAdd(a, b byte) byte {
	return a ^ b
}

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
