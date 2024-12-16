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
