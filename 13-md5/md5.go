package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	hash := md5([]byte("a"))
	fmt.Printf("%x\n", hash)
	// Should print: 0cc175b9c0f1b6a831c399e269772661
}

func md5(message []byte) [16]byte {
	// Initialize variables
	var a0 uint32 = 0x67452301
	var b0 uint32 = 0xefcdab89
	var c0 uint32 = 0x98badcfe
	var d0 uint32 = 0x10325476

	// Calculate padded message length
	// It needs to be congruent to 56 (mod 64)
	messageLen := len(message)
	paddedLen := messageLen + 1 // Add 1 for the 0x80 byte
	if paddedLen%64 > 56 {
		paddedLen += 64 - (paddedLen % 64) + 56
	} else if paddedLen%64 < 56 {
		paddedLen += 56 - (paddedLen % 64)
	}
	paddedLen += 8 // Add 8 bytes for the length

	// Create padded message
	paddedMessage := make([]byte, paddedLen)
	copy(paddedMessage, message)
	paddedMessage[messageLen] = 0x80

	// Append original length in bits at the end
	messageLenBits := uint64(messageLen * 8)
	binary.LittleEndian.PutUint64(paddedMessage[paddedLen-8:], messageLenBits)

	// Constants
	var K [64]uint32
	for i := 0; i < 64; i++ {
		K[i] = uint32(math.Floor(math.Abs(math.Sin(float64(i+1))) * (1 << 32)))
	}

	s := [...]uint{
		7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
		5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
		4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
		6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
	}

	// Process each 512-bit chunk
	numChunks := len(paddedMessage) / 64
	for chunk := 0; chunk < numChunks; chunk++ {
		// Break chunk into sixteen 32-bit words
		M := make([]uint32, 16)
		for j := 0; j < 16; j++ {
			offset := chunk*64 + j*4
			M[j] = binary.LittleEndian.Uint32(paddedMessage[offset : offset+4])
		}

		A := a0
		B := b0
		C := c0
		D := d0

		// Main loop
		for i := 0; i < 64; i++ {
			var F uint32
			var g int

			switch {
			case i < 16:
				F = (B & C) | ((^B) & D)
				g = i
			case i < 32:
				F = (D & B) | ((^D) & C)
				g = (5*i + 1) % 16
			case i < 48:
				F = B ^ C ^ D
				g = (3*i + 5) % 16
			default:
				F = C ^ (B | (^D))
				g = (7 * i) % 16
			}

			F = F + A + K[i] + M[g]
			A = D
			D = C
			C = B
			B = B + leftShift(F, s[i])
		}

		a0 += A
		b0 += B
		c0 += C
		d0 += D
	}

	var digest [16]byte
	binary.LittleEndian.PutUint32(digest[0:], a0)
	binary.LittleEndian.PutUint32(digest[4:], b0)
	binary.LittleEndian.PutUint32(digest[8:], c0)
	binary.LittleEndian.PutUint32(digest[12:], d0)

	return digest
}

func leftShift(x uint32, c uint) uint32 {
	return (x << c) | (x >> (32 - c))
}

func swapUint64(val uint64) uint64 {
    val = ((val << 8) & 0xFF00FF00FF00FF00) | ((val >> 8) & 0x00FF00FF00FF00FF)
    val = ((val << 16) & 0xFFFF0000FFFF0000) | ((val >> 16) & 0x0000FFFF0000FFFF)

    return (val << 32) | (val >> 32)
}
