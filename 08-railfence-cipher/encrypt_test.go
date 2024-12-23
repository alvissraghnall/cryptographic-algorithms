package main

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name       string
		plaintext  string
		rails      int
		wantCiphertext string
		wantErr      bool
	}{
		{
			name:       "Simple encryption",
			plaintext:  "WEAREDISCOVEREDSAVEYOURSELF",
			rails:      3,
			wantCiphertext: "WECRAOEERDSOEESVYUSLAIVDERF",
			wantErr:      false,
		},
		{
			name:       "Encryption with multiple words",
			plaintext:  "ATTACKATDAWN",
			rails:      6,
			wantCiphertext: "AWTANTDATCAK",
			wantErr:      false,
		},
		{
			name:       "Encryption with single word",
			plaintext:  "HELLO",
			rails:      4,
			wantCiphertext: "HELOL",
			wantErr:      false,
		},
		{
			name:       "Encryption with single character",
			plaintext:  "A",
			rails:      3,
			wantCiphertext: "A",
			wantErr:      false,
		},
		{
			name:       "Invalid plaintext",
			plaintext:  "",
			rails:      3,
			wantCiphertext: "",
			wantErr:      true,
		},
		{
			name:       "Invalid rails",
			plaintext:  "HELLO",
			rails:      0,
			wantCiphertext: "",
			wantErr:      true,
		},
		{
			name:       "Invalid rails (negative)",
			plaintext:  "HELLO",
			rails:      -3,
			wantCiphertext: "",
			wantErr:      true,
		},
                {
                        name:       "Plaintext with spaces and lowercase letters",
                        plaintext:  "HElLo FROM iTaLy",
                        rails:      7,
                        wantCiphertext: "HLEAYLTLIOMFOR",
                        wantErr:      false,
                },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCiphertext, err := Encrypt(tt.plaintext, tt.rails)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCiphertext != tt.wantCiphertext {
				t.Errorf("Encrypt() = %v, want %v", gotCiphertext, tt.wantCiphertext)
			}
		})
	}
}
