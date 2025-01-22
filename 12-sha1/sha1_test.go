package main

import (
	"encoding/hex"
	"reflect"
	"testing"
	sha "crypto/sha1"
)

func TestSHA1(t *testing.T) {
	h := sha.New()
	h.Write(bytes65())

	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty string",
			input:    []byte(""),
			expected: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		{
			name:     "hello",
			input:    []byte("hello"),
			expected: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
		{
			name:     "hello world",
			input:    []byte("hello world"),
			expected: "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
		},
		{
			name:     "numeric input",
			input:    []byte("12345"),
			expected: "8cb2237d0679ca88db6464eac60da96345513964",
		},
		{
			name:     "special characters",
			input:    []byte("!@#$%^&*()"),
			expected: "bf24d65c9bb05b9b814a966940bcfa50767c8a8d",
		},
		{
			name:     "long string",
			input:    []byte("The quick brown fox jumps over the lazy dog"),
			expected: "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12",
		},
		{
			name:     "unicode characters",
			input:    []byte("Hello, 世界"),
			expected: "ec105952aaab47ed409894bea51b26b641361df7",
		},
		{
			name:     "repeating characters",
			input:    []byte("aaaaaaaaaa"),
			expected: "3495ff69d34671d1e15b33a63c1379fdedd3a32a",
		},
		{
			name:     "binary data",
			input:    []byte{0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF},
			expected: "aa7432728cff697c0ff81508daf775ac8bf97b08", 
		},
		{
			name:     "64 bytes (block size)",
			input:    bytes64(),
			expected: "c6138d514ffa2135bfce0ed0b8fac65669917ec7",
		},
		{
			name:     "65 bytes (block size + 1)",
			input:    bytes65(),
//			expected: "69bd728ad6e13cd76ff19751fde427b00e395746",
			expected: hex.EncodeToString(h.Sum(nil)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sha1(tt.input)
			resultHex := hex.EncodeToString(result[:])
			if resultHex != tt.expected {
				t.Errorf("sha1(%q) = %s; want %s", tt.input, resultHex, tt.expected)
			}
		})
	}
}

// TestSHA1Consistency ensures that multiple calls with the same input produce the same output
func TestSHA1Consistency(t *testing.T) {
	input := []byte("test consistency")
	first := sha1(input)
	second := sha1(input)

	if !reflect.DeepEqual(first, second) {
		t.Errorf("Inconsistent results for same input:\nFirst:  %x\nSecond: %x", first, second)
	}
}

// TestSHA1DoesNotModifyInput ensures that the input slice is not modified
func TestSHA1DoesNotModifyInput(t *testing.T) {
	input := []byte("test input modification")
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)

	sha1(input)

	if !reflect.DeepEqual(input, inputCopy) {
		t.Errorf("Input was modified during SHA1 calculation:\nOriginal: %v\nModified: %v", inputCopy, input)
	}
}

// TestSHA1WithNilInput tests behavior with nil input
func TestSHA1WithNilInput(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("sha1(nil) panicked: %v", r)
		}
	}()

	result := sha1(nil)
	expected := "da39a3ee5e6b4b0d3255bfef95601890afd80709" // SHA1 of empty string
	resultHex := hex.EncodeToString(result[:])
	if resultHex != expected {
		t.Errorf("sha1(nil) = %s; want %s", resultHex, expected)
	}
}

// Helper functions to generate test data
func bytes64() []byte {
	result := make([]byte, 64)
	for i := range result {
		result[i] = byte(i)
	}
	return result
}

func bytes65() []byte {
	result := make([]byte, 65)
	for i := range result {
		result[i] = byte(i)
	}
	return result
}

// Benchmark SHA1 performance
func BenchmarkSHA1(b *testing.B) {
	input := []byte("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sha1(input)
	}
}

// Benchmark SHA1 with large input
func BenchmarkSHA1Large(b *testing.B) {
	input := make([]byte, 1024*1024) // 1MB
	for i := range input {
		input[i] = byte(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sha1(input)
	}
}
