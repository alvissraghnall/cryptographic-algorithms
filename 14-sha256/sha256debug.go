package main

import (
	sha256builtin "crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

// debugSHA256 compares My SHA-256 implementation with the standard library's implementation
func debugSHA256() {
	// Test vectors from NIST
	testCases := []struct {
		input    string
		expected string
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{"abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
			"248d6a61d20638b8e5c026930c3e6039a33ce45964ff2167f6ecedd419db06c1"},
	}

	fmt.Println("SHA-256 Debug Tool")
	fmt.Println("=================")

	for i, tc := range testCases {
		fmt.Printf("\nTest Case %d:\n", i+1)
		fmt.Printf("Input: %q\n", tc.input)

		// Calculate standard SHA-256
		standardHash := sha256builtin.New()
		io.WriteString(standardHash, tc.input)
		standardResult := hex.EncodeToString(standardHash.Sum(nil))

		fmt.Println("\nExpected hash:", tc.expected)
		fmt.Println("Standard lib:", standardResult)

		// Calculate My SHA-256 implementation
		myResult, intermediateStates := sha256([]byte(tc.input))
		fmt.Println("My implementation:", hex.EncodeToString(myResult[:]))

		// Print intermediate states
		fmt.Println("\nIntermediate States:")
		for i, state := range intermediateStates {
			fmt.Printf("Block %d: %08x %08x %08x %08x %08x %08x %08x %08x\n",
				i+1, state[0], state[1], state[2], state[3], state[4], state[5], state[6], state[7])
		}

		fmt.Println(strings.Repeat("-", 50))
	}
}
