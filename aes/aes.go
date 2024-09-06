package aes


func generate_round_keys(secret_key int) {

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

	temp[0]   = state[4]
	state[4]  = state[5]
	state[5]  = state[6]
	state[6]  = state[7]
	state[7]  = temp[0]

	temp[0]   = state[8]
	temp[1]   = state[9]
	state[8]  = state[10]
	state[9]  = state[11]
	state[10] = temp[0]
	state[11] = temp[1]

	temp[0]   = state[12]
	temp[1]   = state[13]
	temp[2]   = state[14]
	state[12] = state[15]
	state[13] = temp[0]
	state[14] = temp[1]
	state[15] = temp[2]

	return state
}

func MixColumn(c0, c1, c2, c3 *byte) {

	temp := [4]byte{ 0, 0, 0, 0 };

	temp[0] = GFMul(0x02, *c0) ^ GFMul(0x03, *c1) ^ *c2 ^ *c3;
	temp[1] = *c0 ^ GFMul(0x02, *c1) ^ GFMul(0x03, *c2) ^ *c3;
	temp[2] = *c0 ^ *c1 ^ GFMul(0x02, *c2) ^ GFMul(0x03, *c3);
	temp[3] = GFMul(0x03, *c0) ^ *c1 ^ *c2 ^ GFMul(0x02, *c3);

	*c0 = temp[0]
	*c1 = temp[1]
	*c2 = temp[2]
	*c3 = temp[3]

}

func MixColumns(state [16]byte) [16] byte {

	for ix := 0; ix < 4; ix++ {
		MixColumn(&state[0 + ix], &state[4 + ix], &state[8 + ix], &state[12 + ix])
	}

	return state
}

func addKey() {

}


func aesEncrypt (state [16]byte, key string) {

	// First get round keys


	// Add round keys

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

		if a & 0x80 != 0 {
			a = (a << 1) ^ 0x1b 
		} else {
			a <<= 1
		}

		b >>= 1
	}
	return result
}

