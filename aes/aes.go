package aes


func generate_round_keys(secret_key int) {

}




// Transformations


func subBytes() {

}


func ShiftRows(state [16]byte) [16]byte {

	// 0 1 2 3
	// 4 5 6 7
	// 8 9 0 1
	// 2 3 4 5
	var temp [3]byte

	temp[0] = state[4]
	state[4] = state[5]
	state[5] = state[6]
	state[6] = state[7]
	state[7] = temp[0]

	temp[0] = state[8]
	temp[1] = state[9]
	state[8] = state[10]
	state[9] = state[11]
	state[10] = temp[0]
	state[11] = temp[1]

	temp[0] = state[12]
	temp[1] = state[13]
	temp[2] = state[14]
	state[12] = state[15]
	state[13] = temp[0]
	state[14] = temp[1]
	state[15] = temp[2]

	return state
}

func mixColumns() {

}

func addKey() {

}


func aesEncrypt (state [16]byte, key string) {

	// First get round keys


	// Add round keys

}


func gfAdd(a, b byte) {
	
}
