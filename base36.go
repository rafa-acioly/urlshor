package main

import "math"

var (
	base36 = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z'}

	index = map[byte]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14,
		'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
		'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24,
		'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
		'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34,
		'Z': 35,
		'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14,
		'f': 15, 'g': 16, 'h': 17, 'i': 18, 'j': 19,
		'k': 20, 'l': 21, 'm': 22, 'n': 23, 'o': 24,
		'p': 25, 'q': 26, 'r': 27, 's': 28, 't': 29,
		'u': 30, 'v': 31, 'w': 32, 'x': 33, 'y': 34,
		'z': 35,
	}
)

// Encode encodes a number to base36
func encode36(value uint64) string {

	var res [16]byte
	var i int
	for i = len(res) - 1; value != 0; i-- {
		res[i] = base36[value%36]
		value /= 36
	}
	return string(res[i+1:])
}

// Decode decodes a base36-encoded string
func decode36(s string) uint64 {

	res := uint64(0)
	l := len(s) - 1
	for idx := range s {
		c := s[l-idx]
		byteOffset := index[c]
		res += uint64(byteOffset) * uint64(math.Pow(36, float64(idx)))
	}
	return res
}
