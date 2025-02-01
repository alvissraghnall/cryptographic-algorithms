package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	setupConstants()

	fmt.Printf("h0 = 0x%08x\n", h0)
	fmt.Printf("h1 = 0x%08x\n", h1)
	fmt.Printf("h2 = 0x%08x\n", h2)
	fmt.Printf("h3 = 0x%08x\n", h3)
	fmt.Printf("h4 = 0x%08x\n", h4)
	fmt.Printf("h5 = 0x%08x\n", h5)
	fmt.Printf("h6 = 0x%08x\n", h6)
	fmt.Printf("h7 = 0x%08x\n", h7)

	fmt.Println("\nFirst few k values:")
	for i := 0; i < 5; i++ {
		fmt.Printf("k[%d] = 0x%08x\n", i, k[i])
	}

  sha256([]byte {0x80})
}

func sha256 (message []byte) {
  messageLen := len(message)

  paddedMessage := bytes.NewBuffer(message)

  // Pre-processing 
  paddedMessage.WriteByte(0x80)

  padLength := 56 - paddedMessage.Len() % 64

  if padLength < 0 {
    padLength += 64
  }
  for i := 0; i < padLength; i++ {
    paddedMessage.WriteByte(0x00)
  }

  messageLengthInBits := uint64(messageLen) * 8

  binary.Write(paddedMessage, binary.BigEndian, messageLengthInBits)

  var block [16]uint32  // 16 * 32 = 512 bits
  for binary.Read(paddedMessage, binary.BigEndian, &block) == nil {
    // Process block
    // Each block[i] contains 4 bytes in big-endian order
      
    var w [64]uint32

    copy(w[:], block[:])

    equal := true
    for i := 0; i < 16; i++ {
      if w[i] != block[i] {
        equal = false
        break
      }
    }

    // Assert the result
    if equal {
      fmt.Println("The first 16 elements of w and block are equal.")
    } else {
      fmt.Println("The first 16 elements of w and block are NOT equal.")
    } 
  } 
}
