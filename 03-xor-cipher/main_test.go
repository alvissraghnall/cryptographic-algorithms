package main

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name       string
		plainText  string
		key         string
		wantCipher  string
		wantErr     bool
	}{
		{"valid input", "HELLO", "101ComputingKey", "", false},
		{"empty plaintext", "", "101ComputingKey", "", false},
		{"empty key", "HELLO", "", "", true},
		{"non-binary key", "HELLO", "123ComputingKey", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Encrypt(tt.plainText, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name         string
		cipherText   string
		key           string
		wantPlain     string
		wantErr       bool
	}{
		{"valid input", "0111100101110101011111010000111100100000", "101ComputingKey", "", false},
		{"invalid binary cipher", "1234567890", "101ComputingKey", "", true},
		{"empty cipher", "", "101ComputingKey", "", false},
		{"empty key", "0111100101110101011111010000111100100000", "", "", true},
		{"non-binary key", "0111100101110101011111010000111100100000", "123ComputingKey", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Decrypt(tt.cipherText, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBinary(t *testing.T) {
	tests := []struct {
		name       string
		s           string
		wantBinary  string
	}{
		{"valid binary", "10101010", "10101010"},
		{"non-binary", "HELLO", "0100100001100101011011000110110001101111"},
		{"empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBinary := binary(tt.s); gotBinary != tt.wantBinary {
				t.Errorf("binary() = %v, want %v", gotBinary, tt.wantBinary)
			}
		})
	}
}

func TestBinaryToString(t *testing.T) {
	tests := []struct {
		name       string
		binary      string
		wantString  string
	}{
		{"valid binary", "0100100001100101011011000110110001101111", "HELLO"},
		{"invalid binary", "1234567890", ""},
		{"empty binary", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotString := binaryToString(tt.binary); gotString != tt.wantString {
				t.Errorf("binaryToString() = %v, want %v", gotString, tt.wantString)
			}
		})
	}
}

func TestPadBinary(t *testing.T) {
	tests := []struct {
		name       string
		binary      string
		wantPadded  string
	}{
		{"valid binary", "1010", "00001010"},
		{"empty binary", "", "00000000"},
		{"already padded binary", "00001010", "00001010"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPadded := padBinary(tt.binary); gotPadded != tt.wantPadded {
				t.Errorf("padBinary() = %v, want %v", gotPadded, tt.wantPadded)
			}
		})
	}
}
