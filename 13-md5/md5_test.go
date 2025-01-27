package main

import (
	"encoding/hex"
	"testing"
	md5builtin "crypto/md5"
)

type md5Test struct {
	input    []byte
	expected string
}

func TestMD5Implementation(t *testing.T) {
	byteArr := []byte{0x80, 0x57, 0x9A, 0x8b, 0x6d, 0x27, 0x08}
	nativeHash := md5builtin.Sum(byteArr)
	tests := []md5Test{
		// Empty string test
		{
			input:    []byte(""),
			expected: "d41d8cd98f00b204e9800998ecf8427e",
		},
		// Basic ASCII strings
		{
			input:    []byte("a"),
			expected: "0cc175b9c0f1b6a831c399e269772661",
		},
		{
			input:    []byte("abc"),
			expected: "900150983cd24fb0d6963f7d28e17f72",
		},
		{
			input:    []byte("message digest"),
			expected: "f96b697d7cb7938d525a2f31aaf161d0",
		},
		// Longer strings
		{
			input:    []byte("abcdefghijklmnopqrstuvwxyz"),
			expected: "c3fcd3d76192e4007dfb496cca67e13b",
		},
		{
			input:    []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
			expected: "d174ab98d277d9f5a5611c2c9f419d9f",
		},
		// Special characters
		{
			input:    []byte("+#+#+"),
			expected: "e8cd544145b2aef906ee12a616230fd2",
		},
		// Large string to test multiple blocks
		{
			input:    []byte("SAM SJSNENWJWJHSBSHSHSHSHSHWBSJAOKAjjwjsjjs722eijejdjsjsjsjsj"),
			expected: "9b678d93bf00493a61e58a6b3f7cfa2d",
		},
		// Binary data
		{
			input:    byteArr,
			expected: hex.EncodeToString(nativeHash[:]),
		},
		// Test padding edge cases
		{
			input:    make([]byte, 55), // exactly 447 bits, requiring minimal padding
			expected: "c9ea3314b91c9fd4e38f9432064fd1f2",
		},
		{
			input:    make([]byte, 56), // exactly 448 bits, requiring new block
			expected: "e3c4dd21a9171fd39d208efa09bf7883",
		},
		{
			input:    make([]byte, 63), // testing near block boundary
			expected: "65cecfb980d72fde57d175d6ec1c3f64",
		},
		{
			input:    make([]byte, 64), // exactly one block
			expected: "3b5d3c7d207e37dceeedd301e35e2e58",
		},
	}

	for i, test := range tests {
		result := md5_buf(test.input)
		got := hex.EncodeToString(result[:])
		if got != test.expected {
			t.Errorf("Test %d failed:\nInput: %q\nExpected: %s\nGot: %s",
				i, string(test.input), test.expected, got)
		}
	}

	for i, test := range tests {
		result := md5(test.input)
		got := hex.EncodeToString(result[:])
		if got != test.expected {
			t.Errorf("Test %d failed:\nInput: %q\nExpected: %s\nGot: %s",
				i, string(test.input), test.expected, got)
		}
	}
}

// Test helper functions
func TestLeftShift(t *testing.T) {
	tests := []struct {
		value    uint32
		shift    uint
		expected uint32
	}{
		{0x12345678, 4, 0x23456781},
		{0xFFFFFFFF, 8, 0xFFFFFFFF},
		{0x00000001, 31, 0x80000000},
		{0x80000000, 1, 0x00000001},
	}

	for i, test := range tests {
		result := leftShift(test.value, test.shift)
		if result != test.expected {
			t.Errorf("Test %d failed: leftShift(0x%x, %d) = 0x%x, expected 0x%x",
				i, test.value, test.shift, result, test.expected)
		}
	}
}

func TestSwapUint64(t *testing.T) {
	tests := []struct {
		input    uint64
		expected uint64
	}{
		{0x1234567890ABCDEF, 0xEFCDAB9078563412},
		{0x0000000000000000, 0x0000000000000000},
		{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
		{0x0000000000000001, 0x0100000000000000},
	}

	for i, test := range tests {
		result := swapUint64(test.input)
		if result != test.expected {
			t.Errorf("Test %d failed: swapUint64(0x%x) = 0x%x, expected 0x%x",
				i, test.input, result, test.expected)
		}
	}
}

// Benchmark tests
func BenchmarkMD5Empty(b *testing.B) {
	input := []byte("")
	for i := 0; i < b.N; i++ {
		md5(input)
	}
}

func BenchmarkMD5Small(b *testing.B) {
	input := []byte("abc")
	for i := 0; i < b.N; i++ {
		md5(input)
	}
}

func BenchmarkMD5Large(b *testing.B) {
	input := make([]byte, 1024*1024) // 1MB
	for i := 0; i < b.N; i++ {
		md5(input)
	}
}
