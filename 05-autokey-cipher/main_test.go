

package main

import (
	"testing"
)

func TestEncipher(t *testing.T) {
	tests := []struct {
		name       string
		plainText  string
		key        string
		want       string
		wantErr    bool
	}{
		{"simple", "HELLO", "n", "ULPWZ", false},
		{"with spaces", "attack at dawn", "queenly", "QNXEPVYTWTWP", false},
		{"invalid plain text", "HELLO!", "KEY", "", true},
		{"invalid key", "HELLO", "KEY!", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encipher(tt.plainText, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encipher() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDecipher(t *testing.T) {
	tests := []struct {
		name       string
		cipherText string
		key        string
		want       string
		wantErr    bool
	}{
		{"simple", "qnxepvytwtwp", "QUeEnlY", "ATTACKATDAWN", false},
		{"with spaces", "yyfvbvugamk", "function", "TESTINGTHIS", false},
		{"invalid cipher text", "encrypted text!", "KEY", "", true},
		{"invalid key", "encrypted text", "KEY!", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decipher(tt.cipherText, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decipher() = %q, want %q", got, tt.want)
			}
		})
	}
}
