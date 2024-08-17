package main

import (
	"bytes"
	"encoding/binary"
//	"fmt"
//	"strings"
	"math"
)

var (
	h0, h1, h2, h3, h4, h5, h6, h7 uint32
	k                               []uint32
)

func sieveOfEratosthenes(n int) []int {
	limit := 313

	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= limit; p++ {
		if isPrime[p] {
			for i := p * p; i <= limit; i += p {
				isPrime[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= limit && len(primes) < n; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func extractFractionalBitsToHex(num int, isCubeRoot bool) uint32 {
	var root float64
	if isCubeRoot {
		root = math.Pow(float64(num), 1.0/3.0)
	} else {
		root = math.Sqrt(float64(num))
	}

	_, frac := math.Modf(root)

	var bits uint32
	for i := 0; i < 32; i++ {
		frac *= 2
		bit := uint32(frac)
		bits = (bits << 1) | bit
		frac -= float64(bit)
	}

	return bits
}

func setupConstants() {
	primes := sieveOfEratosthenes(64)
	k = make([]uint32, 64)

	// Extract and store values
	for i, prime := range primes {
		// Square root hex values (first 8 primes)
		if i < 8 {
			sqrtHexValue := extractFractionalBitsToHex(prime, false)
			switch i {
			case 0:
				h0 = sqrtHexValue
			case 1:
				h1 = sqrtHexValue
			case 2:
				h2 = sqrtHexValue
			case 3:
				h3 = sqrtHexValue
			case 4:
				h4 = sqrtHexValue
			case 5:
				h5 = sqrtHexValue
			case 6:
				h6 = sqrtHexValue
			case 7:
				h7 = sqrtHexValue
			}
		}

		// Cube root hex values
		k[i] = extractFractionalBitsToHex(prime, true)
	}
}

func main() {

}

func sha256 (message []byte) ([32]byte, [][8]uint32) {
  
  setupConstants()

  messageLen := len(message)

  paddedMessage := bytes.NewBuffer(message)

  // Pre-processing 
  paddedMessage.WriteByte(0x80)

  padLength := 55 - messageLen % 64

  if padLength < 0 {
    padLength += 64
  }
  for i := 0; i < padLength; i++ {
    paddedMessage.WriteByte(0x00)
  }

  messageLengthInBits := uint64(messageLen) * 8

  binary.Write(paddedMessage, binary.BigEndian, messageLengthInBits)

  var block [16]uint32  // 16 * 32 = 512 bits
  var intermediateStates [][8]uint32

  for binary.Read(paddedMessage, binary.BigEndian, &block) == nil {
    // Process block
    // Each block[i] contains 4 bytes in big-endian order
      
    var w [64]uint32

    copy(w[:], block[:])

    var s0, s1 uint32
    for i := 16; i < 64; i++ {
      s0 = (rightRotate32(w[i - 15], 7)) ^ (rightRotate32(w[i - 15], 18)) ^ (w[i - 15] >> 3)
      s1 = (rightRotate32(w[i - 2], 17)) ^ (rightRotate32(w[i - 2], 19)) ^ (w[i - 2] >> 10)

      w[i] = w[i - 16] + s0 + w[i - 7] + s1

    }

    // Initialize working variables to current hash value:
    a, b, c, d, e, f, g, h := h0, h1, h2, h3, h4, h5, h6, h7

    for j := range 64 {
      S1 := ((rightRotate32(e, 6)) ^
        (rightRotate32(e, 11)) ^ 
        rightRotate32(e, 25))

      ch := (e & f) ^ ((^e) & g)
      temp1 := h + S1 + ch + k[j] + w[j]
      S0 := ((rightRotate32(a, 2)) ^
        (rightRotate32(a, 13)) ^
        (rightRotate32(a, 22)))
      maj := (a & b) ^ (a & c) ^ (b & c)
      temp2 := S0 + maj
      
      h, g, f, e, d, c, b, a = g, f, e, d+temp1, c, b, a, temp1+temp2
    }

    // Add the compressed chunk to the current hash value:
    h0, h1, h2, h3, h4, h5, h6, h7 = h0+a, h1+b, h2+c, h3+d, h4+e, h5+f, h6+g, h7+h
	
	intermediateStates = append(intermediateStates, [8]uint32{h0, h1, h2, h3, h4, h5, h6, h7})
  } 

  var digest bytes.Buffer

  // Write all uint32 values (h0 to h7) to the buffer in big-endian format
  binary.Write(&digest, binary.BigEndian, []uint32{h0, h1, h2, h3, h4, h5, h6, h7})

  digestByteArr := digest.Bytes()

  // fmt.Printf("%x\n\n", digestByteArr)
  return [32]byte(digestByteArr), intermediateStates

}

func rightRotate32(value uint32, shift uint) uint32 {
	shift &= 0x1F // Mask to ensure shift is within 0-31
	return (value >> shift) | (value << (32 - shift))
}
