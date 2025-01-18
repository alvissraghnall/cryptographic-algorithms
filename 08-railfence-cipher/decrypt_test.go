
package main

import (
	"testing"
)

// TestDecryptRailFence is the main test function for DecryptRailFence.
func TestDecryptRailFence(t *testing.T) {
	tests := []struct {
		name        string
		ciphertext  string
		rails       int
		expected    string
		expectError bool
	}{
		// Basic test cases
		{"Basic-3-Rails", "WECRLTEERDSOEEFEAOCAIVDEN", 3, "WEAREDISCOVEREDFLEEATONCE", false},
		{"Basic-4-Rails", "WIREEEDSEEEACAECVDLTNROFO", 4, "WEAREDISCOVEREDFLEEATONCE", false},
		// Edge cases
		{"Single-Rail", "WEAREDISCOVEREDFLEEATONCE", 1, "WEAREDISCOVEREDFLEEATONCE", false},
		{"Two-Rails", "WAEICVRDLETNEERDSOEEFEAOC", 2, "WEAREDISCOVEREDFLEEATONCE", false},
		{"Empty-Input", "", 3, "Invalid ciphertext, and/or rail length provided.", true},
		{"One-Character", "A", 3, "A", false},
		// Invalid input scenarios
		{"Zero-Rails", "ANYTEXT", 0, "Invalid ciphertext, and/or rail length provided.", true},
		{"Negative-Rails", "ANYTEXT", -3, "Invalid ciphertext, and/or rail length provided.", true},
		// Larger inputs
		{"Large-Text-3-Rails", "TIEIXLHSSTSRNEAPEIATGM", 3, "THISISATESTRINGEXAMPLE", false},
		{"Large-Text-4-Rails", "TANPHSTIGMLIIETEAESSX", 4, "THISISATESTINGEXAMPLE", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DecryptRailFence(tt.ciphertext, tt.rails)
			if err != nil && !tt.expectError {
				t.Errorf("DecryptRailFence(%q, %d) returned error %q; want no error", tt.ciphertext, tt.rails, err)
				return
			}
			if err == nil && tt.expectError {
				t.Errorf("DecryptRailFence(%q, %d) did not return error; want error", tt.ciphertext, tt.rails)
				return
			}
			if tt.expectError {
				if err.Error() != tt.expected {
					t.Errorf("DecryptRailFence(%q, %d) = %q; want %q", tt.ciphertext, tt.rails, err.Error(), tt.expected)
				}
			} else {
				if result != tt.expected {
					t.Errorf("DecryptRailFence(%q, %d) = %q; want %q", tt.ciphertext, tt.rails, result, tt.expected)
				}
			}
		})
	}
}

// BenchmarkDecryptRailFence benchmarks the DecryptRailFence function.
func BenchmarkDecryptRailFence(b *testing.B) {
	ciphertext := "WECRLTEERDSOEEFEAOCAIVDEN"
	rails := 3
	for i := 0; i < b.N; i++ {
		DecryptRailFence(ciphertext, rails)
	}
}
