package main

import (
	builtin "crypto/sha256"
	"encoding/hex"
	"testing"
)

// TestSHA256 compares the output of your SHA-256 implementation with the standard library's implementation.
func TestSHA256(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		// NIST CAVP Test Vectors
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"}, // Empty string
		{"abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"}, // "abc"
		{"abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq", "248d6a61d20638b8e5c026930c3e6039a33ce45964ff2167f6ecedd419db06c1"}, // 56-byte string
		{"abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu", "cf5b16a778af8380036ce59e7b0492370b249b11e8f07a51afac45037afee9d1"}, // 112-byte string

		// Edge Cases
		{"a", "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"}, // Single character
		{"1234567890", "f371bc4a311f2b009eef952dd83ca80e2b60026c8e935592d0f9c308453c813e"}, // Numeric string
		{"The quick brown fox jumps over the lazy dog", "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592"}, // Common phrase
		{"The quick brown fox jumps over the lazy dog.", "ef537f25c895bfa782526529a9b63d97aa631564d5d789c2b765448c8635fb6c"}, // Common phrase with a period

		// Custom Inputs
		{"hello world", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"}, // "hello world"
		{"Go is awesome!", "c3a5f4a5b5e5f5d5c5b5a5f4a5b5e5f5d5c5b5a5f4a5b5e5f5d5c5b5a5f4a5b5"}, // Custom input
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			// Compute hash using your implementation
			hash, _ := sha256([]byte(tc.input))
			hashHex := hex.EncodeToString(hash[:])

			// Compute hash using the standard library
			expectedHash := builtin.Sum256([]byte(tc.input))
			expectedHashHex := hex.EncodeToString(expectedHash[:])

			// Compare the results
			if hashHex != expectedHashHex {
				t.Errorf("Input: %s\nExpected: %s\nActual:   %s\n", tc.input, expectedHashHex, hashHex)
			}
		})
	}
}

// TestSHA256IntermediateStates verifies the intermediate states of the SHA-256 computation.
func TestSHA256IntermediateStates(t *testing.T) {
	testCases := []struct {
		input    string
		expected [][8]uint32
	}{
		{
			input: "abc",
			expected: [][8]uint32{
				{0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19}, // Initial state
				{0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19}, // After processing first block
				{0xba7816bf, 0x8f01cfea, 0x414140de, 0x5dae2223, 0xb00361a3, 0x96177a9c, 0xb410ff61, 0xf20015ad}, // Final state
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			_, intermediateStates := sha256([]byte(tc.input))

			// Compare intermediate states
			for i, state := range intermediateStates {
				if state != tc.expected[i] {
					t.Errorf("Block %d:\nExpected: %08x %08x %08x %08x %08x %08x %08x %08x\nActual:   %08x %08x %08x %08x %08x %08x %08x %08x\n",
						i+1,
						tc.expected[i][0], tc.expected[i][1], tc.expected[i][2], tc.expected[i][3],
						tc.expected[i][4], tc.expected[i][5], tc.expected[i][6], tc.expected[i][7],
						state[0], state[1], state[2], state[3], state[4], state[5], state[6], state[7],
					)
				}
			}
		})
	}
}

// TestSHA256Padding verifies the padding logic of the SHA-256 implementation.
func TestSHA256Padding(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"}, // Empty string
		{"a", "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"}, // Single character
		{"1234567890", "f371bc4a311f2b009eef952dd83ca80e2b60026c8e935592d0f9c308453c813e"}, // Numeric string
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			hash, _ := sha256([]byte(tc.input))
			hashHex := hex.EncodeToString(hash[:])

			if hashHex != tc.expected {
				t.Errorf("Input: %s\nExpected: %s\nActual:   %s\n", tc.input, tc.expected, hashHex)
			}
		})
	}
}
