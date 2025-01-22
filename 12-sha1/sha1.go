package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")

	message := []byte("hello")
	sha1(message)
//	x := sha1(message)

//	fmt.Printf("%x", x)
}

func sha1(message []byte) [20]byte {
	var h0, h1, h2, h3, h4 uint32 = 0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476, 0xC3D2E1F0

	ml := len(message) * 8

	// Pre-processing:
	// append the bit '1' to the message e.g. by adding 0x80 if message length is a multiple of 8 bits.
	messageBytes := append([]byte{}, message...)
	messageBytes = append(messageBytes, 0x80)

	// append zero bits (0x00) to message, until messagelen % 512 == 448
	padLength := 56 - (len(messageBytes) % 64)

	if padLength < 0 {
		padLength += 64
	}
	for i := 0; i < padLength; i++ {
		messageBytes = append(messageBytes, 0x00)
	}

	// append ml, the original message length in bits, as a 64-bit big-endian integer.
	mlBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(mlBytes, uint64(ml))
	messageBytes = append(messageBytes, mlBytes...)
	//fmt.Println(len(messageBytes), messageBytes, ml, mlBytes)

	for i := 0; i < len(messageBytes); i += 64 {
		chunk := messageBytes[i:min(i+64, len(messageBytes))]

		w := make([]uint32, 80)

		for i := 0; i < 16; i++ {
			w[i] = (uint32(chunk[i*4]) << 24) | (uint32(chunk[i*4+1]) << 16) | (uint32(chunk[i*4+2]) << 8) | uint32(chunk[i*4+3])
		}

		for k := 16; k < 80; k++ {
			w[k] = leftShift(w[k-3]^w[k-8]^w[k-14]^w[k-16], 1)
		}

		a, b, c, d, e := h0, h1, h2, h3, h4

		for i := range 80 {
			var K, F uint32
			switch {
			case i >= 0 && i <= 19:
				F = (b & c) | (^b & d)
				K = 0x5A827999

			case i >= 20 && i <= 39:
				F = b ^ c ^ d
				K = 0x6ED9EBA1

			case i >= 40 && i <= 59:
				F = (b & c) | (b & d) | (c & d)
				K = 0x8F1BBCDC

			case i >= 60 && i <= 79:
				F = b ^ c ^ d
				K = 0xCA62C1D6
			}

			temp := (leftShift(a, 5)) + F + e + K + w[i]

			e = d
			d = c
			c = leftShift(b, 30)
			b = a
			a = temp
		}

		//Add this chunk's hash to result so far:
		h0 = h0 + a
		h1 = h1 + b
		h2 = h2 + c
		h3 = h3 + d
		h4 = h4 + e

		//fmt.Println(len(w), w)
	}

	var result [20]byte
	hs := [...]uint32{h0, h1, h2, h3, h4}
	for i, h := range hs {
		result[i*4] = byte(h >> 24)
		result[i*4+1] = byte(h >> 16)
		result[i*4+2] = byte(h >> 8)
		result[i*4+3] = byte(h)
	}

	return result
}

func leftShift(value uint32, shift int) uint32 {
	return ((value << shift) | (value >> (32 - shift))) & 0xFFFFFFFF
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
