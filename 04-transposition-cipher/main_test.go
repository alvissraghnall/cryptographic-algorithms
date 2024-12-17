

package main

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name       string
		plainText  string
		key        string
		want       string
		wantErr    bool
	}{
		{"simple", "HELLO", "ABC", "", false},
		{"with spaces", "HELLO WORLD", "ABCDEF", "", false},
		{"with multiple rows", "HELLO\nWORLD", "ABC", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Encrypt(tt.plainText, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEncryptError(t *testing.T) {
	tests := []struct {
		name       string
		plainText  string
		key        string
		wantErr    bool
	}{
		{"invalid plain text", "HELLO!", "ABC", true},
		{"invalid key", "HELLO", "ABC!", true},
		{"empty plain text", "", "ABC", true},
		{"empty key", "HELLO", "", true},
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
