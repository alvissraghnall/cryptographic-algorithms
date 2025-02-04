package main

import (
	builtin "crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

// TestSHA256 compares the output of your SHA-256 implementation with the standard library's implementation.
func TestSHA256(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		// NIST CAVP Test Vectors
		{"empty_string", "", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"abc", "abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{"56_bytes", "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq", "248d6a61d20638b8e5c026930c3e6039a33ce45964ff2167f6ecedd419db06c1"},
		{"112_bytes", "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu", "cf5b16a778af8380036ce59e7b0492370b249b11e8f07a51afac45037afee9d1"},

		// Edge Cases
		{"single_char", "a", "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"},
		{"numbers", "1234567890", "f371bc4a311f2b009eef952dd83ca80e2b60026c8e935592d0f9c308453c813e"},
		{"fox_no_punct", "The quick brown fox jumps over the lazy dog", "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592"},
		{"fox_with_punct", "The quick brown fox jumps over the lazy dog.", "ef537f25c895bfa782526529a9b63d97aa631564d5d789c2b765448c8635fb6c"},

		// Additional Test Cases
		{"all_spaces", "    ", "4e959a1673c71d516d49d23a82f34c6f6159b3f4642cd674f5ff9f7d3e6e3dad"},
		{"special_chars", "!@#$%^&*()", "a620c5c5c85e0001b6b4d55653224193ad8d5c3c70b1132e49af13b76599d46d"},
		{"unicode", "Hello, 世界", "45c71bca470f3634e299214d088c3dd9d92cf54b06a4c625285cc58bfafbcc2a"},
		{"repeated_char", strings.Repeat("a", 1000), "aae4f3bac727b7b4e84baddb9e81d16e955f37579e47962588d80749c46dd1ed"},
		{"binary_data", string([]byte{0x00, 0xFF, 0x00, 0xFF}), "4b45e1bec21185865d2b6b7c44f6f04c2896c9936edd6d51c3c5e89e7ee01b97"},
		
		// Block Boundary Tests
		{"63_bytes", strings.Repeat("x", 63), "46150c8b234e2f5af71e5b8e3924188601e87592d12a56f5ee538cb6fd93494d"},
		{"64_bytes", strings.Repeat("x", 64), "23097886e67f06c966d0a8df1c6e6c10f5c76da5339070f0689a35509645d56f"},
		{"65_bytes", strings.Repeat("x", 65), "442ef810344657d5dd9aa1944e3d4c9742019007a17548994e29644c5d4f0dcd"},
		
		// Multiple Block Tests
		{"120_bytes", strings.Repeat("y", 120), "a5f6e457394d665759e8eb8e7cc33586e54cfe43f13c5a3fbc3c1131c39df26c"},
		{"128_bytes", strings.Repeat("y", 128), "47e156d8a4e52a36bcf21304467e8fb3783b12b05b01c9ed8927fea43906953b"},
		{"129_bytes", strings.Repeat("y", 129), "2a63f022baf47a44c5f9ceef8e234d2b85e97fd4d5d9b9dd79c3b389b8f24293"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Compute hash using implementation
			hash, intermediateStates := sha256([]byte(tc.input))
			hashHex := hex.EncodeToString(hash[:])

			// Compute hash using the standard library
			expectedHash := builtin.Sum256([]byte(tc.input))
			expectedHashHex := hex.EncodeToString(expectedHash[:])

			// Compare the results
			if hashHex != expectedHashHex {
				t.Errorf("\nInput: %s\nExpected: %s\nActual:   %s", tc.input, expectedHashHex, hashHex)
			}

			// Log intermediate states if test fails
			if t.Failed() {
				t.Logf("Intermediate states for %s:", tc.name)
				for i, state := range intermediateStates {
					t.Logf("Block %d: %v", i, state)
				}
			}
		})
	}
}

// BenchmarkSHA256 provides comprehensive benchmarking for the SHA-256 implementation
func BenchmarkSHA256(b *testing.B) {
	benchCases := []struct {
		name string
		size int
	}{
		{"Empty", 0},
		{"64B", 64},    // Single block
		{"128B", 128},  // Two blocks
		{"256B", 256},  // Four blocks
		{"512B", 512},  // Eight blocks
		{"1KB", 1024},
		{"4KB", 4096},
		{"1MB", 1024 * 1024},
	}

	for _, bc := range benchCases {
		input := make([]byte, bc.size)
		for i := range input {
			input[i] = byte(i % 256) // Fill with repeating pattern
		}

		b.Run(fmt.Sprintf("Custom_%s", bc.name), func(b *testing.B) {
			b.SetBytes(int64(bc.size))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hash, states := sha256(input)
				_ = hash
				_ = states
			}
		})

		b.Run(fmt.Sprintf("Builtin_%s", bc.name), func(b *testing.B) {
			b.SetBytes(int64(bc.size))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hash := builtin.Sum256(input)
				_ = hash
			}
		})
	}
}

// BenchmarkSHA256Parallel tests parallel performance
func BenchmarkSHA256Parallel(b *testing.B) {
	input := make([]byte, 1024) // 1KB input
	for i := range input {
		input[i] = byte(i % 256)
	}

	b.Run("Custom_Parallel", func(b *testing.B) {
		b.SetBytes(int64(len(input)))
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				hash, states := sha256(input)
				_ = hash
				_ = states
			}
		})
	})

	b.Run("Builtin_Parallel", func(b *testing.B) {
		b.SetBytes(int64(len(input)))
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				hash := builtin.Sum256(input)
				_ = hash
			}
		})
	})
}
