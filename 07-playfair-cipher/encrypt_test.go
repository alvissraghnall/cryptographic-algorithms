package main

import (
	"testing"
)

func TestEncryptPlayfair(t *testing.T) {
	tests := []struct {
		plaintext string
		key       string
		expected  string
		err       bool
	}{
		// Basic tests
		{"DIAMOND", "BIRMINGHAM", "HNCRUDAZ", false},
		{"GENERE", "JAZZ", "HFLGQF", false},

		// Tests with spaces in plaintext
		{"HELLO WORLD", "contributor", "KFPVMCOBCSUZ", false},

		// Tests with repeated letters
		{"BALLOON", "KEY", "KBNVMIIO", false},

		{"communicate", "computer", "OMRMPCSGPTER", false},

		// Tests with J treated as I
		{"JIGSAW PUZZLE", "SECRET", "KWMBAHXOZTYMCW", false},

		// Tests with edge cases
		{"", "SECRET", "", false},       // Empty plaintext
		{"HELLO", "", "", true},        // Empty key
		{"HELLO", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "", true}, // Key > 25 chars

		// Long plaintext
		{"COMSEC means communications security", "galois", "DLFDSDNDIHBDDTNTUEBLUOIMCVBSERULYO", false},
		{ "Hear me and rejoice You have had the privilege of being saved by the great Thanos You may think this is suffering no It is salvation The universal scale tips toward balance because of your sacrifice Smile For even in death you have become children of Thanos", "Thanos", "OBNQUSNOKYCLALDSZNPOTXBONCHABUQKXFULLBTLCSKAFBTXSEDWHABLUDNHHANOTEZNMPNXHAKAFNAGCFCVEMIVLSQKHKOTFAFCCTFZNHLAOHOBROFXDUCTFEICULAFMBHTXHYKCHIOADSCSDOQBSTLZNMUCTDQKGQISBQFULLTUDZSAKDKCOHAZNPOTXSCSDTUSDAGKEUDOTMSANOTCV", false },
		// Non-alphabet characters
		{"HELLO123", "MONARCHY", "", true}, // Invalid characters in plaintext
	}

	for _, tt := range tests {
		t.Run(tt.plaintext+"_"+tt.key, func(t *testing.T) {
			ciphertext, err := encryptPlayfair(tt.plaintext, tt.key)
			if tt.err {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if ciphertext != tt.expected {
					t.Errorf("Expected: %s, Got: %s", tt.expected, ciphertext)
				}
			}
		})
	}
}
