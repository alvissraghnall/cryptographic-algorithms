package main

import (
        "testing"
)

func TestGobEncode(t *testing.T) {
	text := "Hello, World!"
	encoded, err := gobEncode(text)
	if err != nil {
		t.Fatal(err)
	}

	if len(encoded) == 0 {
		t.Errorf("Expected encoded data, got empty slice")
	}
}

func TestGobDecode(t *testing.T) {
	text := "Hello, World!"
	encoded, err := gobEncode(text)
	if err != nil {
		t.Fatal(err)
	}

	var decoded string
	err = gobDecode(encoded, &decoded)
	if err != nil {
		t.Fatal(err)
	}

	if decoded != text {
		t.Errorf("Expected decoded text to match original text")
	}
}

func TestXorCipher(t *testing.T) {
	text := "Hello, World!"
	key := "secretkey"

	cipherText, err := xorCipher(text, key)
	if err != nil {
		t.Fatal(err)
	}

	if len(cipherText) == 0 {
		t.Errorf("Expected cipher text, got empty slice")
	}
}

func TestXorDecipher(t *testing.T) {
	text := "Hello, World!"
	key := "secretkey"

	cipherText, err := xorCipher(text, key)
	if err != nil {
		t.Fatal(err)
	}

	decipheredText, err := xorDecipher(cipherText, key)
	if err != nil {
		t.Fatal(err)
	}

	if decipheredText != text {
		t.Errorf("Expected deciphered text to match original text")
	}
}

func TestXorCipherDecipherRoundTrip(t *testing.T) {
	text := "Hello, World!"
	key := "secretkey"

	cipherText, err := xorCipher(text, key)
	if err != nil {
		t.Fatal(err)
	}

	decipheredText, err := xorDecipher(cipherText, key)
	if err != nil {
		t.Fatal(err)
	}

	if decipheredText != text {
		t.Errorf("Expected deciphered text to match original text")
	}
}
