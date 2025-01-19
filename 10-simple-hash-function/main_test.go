
package main

import (
	"testing"
)

func TestHashAscii(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty string", "", 0},
		{"single character", "a", 97 % 1024},
		{"multiple characters", "hello", (104+101+108+108+111) % 1024},
		{"long string", "this is a very long string", (116+104+105+115+105+115+97+118+101+114+121+108+111+110+103+115+116+114+105+110+103) % 1024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashAscii(tt.input); got != tt.want {
				t.Errorf("hashAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashBinaryShift(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty string", "", 0},
		{"single character", "a", (0<<5+97) % 1024},
		{"multiple characters", "hello", ((((((0<<5)+104)<<5+101)<<5+108)<<5+108)<<5+111) % 1024},
		{"long string", "this is a very long string", (((((((((((((((((((((((0<<5)+116)<<5)+104)<<5+105)<<5+115)<<5+105)<<5+115)<<5+97)<<5+118)<<5+101)<<5+114)<<5+121)<<5+108)<<5+111)<<5+110)<<5+103)<<5+115)<<5+116)<<5+114)<<5+105)<<5+110)<<5+103) % 1024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashBinaryShift(tt.input); got != tt.want {
				t.Errorf("hashBinaryShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashFunctionsCollision(t *testing.T) {
	// Test for collision between hashAscii and hashBinaryShift
	inputs := []string{"hello", "world", "this is a test"}
	for _, input := range inputs {
		asciiHash := hashAscii(input)
		binaryShiftHash := hashBinaryShift(input)
		if asciiHash == binaryShiftHash {
			t.Errorf("Collision detected between hashAscii and hashBinaryShift for input %q", input)
		}
	}
}
