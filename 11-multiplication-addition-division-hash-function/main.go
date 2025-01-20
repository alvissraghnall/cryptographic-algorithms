package main

import (
	"fmt"
)

func main() {
	stung := fnv1a("STUNG")
	hex := fmt.Sprintf("%x", stung)

	fmt.Println(hex)
}

func fnv1a (key string) uint64 {
	var FNV_OFFSET_BASIS uint64 = 14695981039346656037
	var FNV_PRIME uint64 = 1099511628211

	hashValue := FNV_OFFSET_BASIS

	for _, byt := range []byte(key) {
		hashValue = (hashValue ^ uint64(byt)) * FNV_PRIME
	}

	return hashValue
}

func isPrime (num uint64) bool {
	if num < 2 { return false }
	if num == 2 || num == 3 {
		return true
	}

	if num % 2 == 0 || num % 3 == 0 { return false }

	i := uint64(5)
	for i * i <= num {
		if num % i == 0 || (num % (i + 2)) == 0 {
			return false
		}
		i += 6
	}

	return true
}

func genNextPrime (n uint64) uint64 {
	if n <= 2 { return 2 }
	var candidate uint64

	if n % 2 != 0 {
		candidate = n
	} else {
		candidate = n + 1
	}
	
	for !isPrime(candidate) {
		candidate += uint64(2)
	}

	return candidate
}

func genMADVariables (fnvHash uint64, hashTableSize int) (a, b uint64, p uint64, m int) {
	//  `a`: odd number
	a = 3

	// `b`: non-negative number
	b = 17

	// `p`: Prime number greater than the max FNV-1a hash value
	p = 18446744073709551557

	// `m`: Hash table size
	m = hashTableSize

	return
}

func madHash(text string) int {
	hashCode := fnv1a(text)
	a, b, p,m := genMADVariables(hashCode, 100)
	return int(((a * hashCode + b) % p) % uint64(m))
}

