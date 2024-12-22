package main

import "testing"

func TestDecrypt(t *testing.T) {
	plainText := "HELP"
	key := "DDCF"
	cipherText, err := Encrypt(plainText, key)
	if err != nil {
		t.Errorf("Encrypt returned error: %v", err)
	}

	decryptedText, err := Decrypt(cipherText, key)
	if err != nil {
		t.Errorf("Decrypt returned error: %v", err)
	}

	if decryptedText != plainText {
		t.Errorf("Decrypted text does not match original plain text: %v != %v", decryptedText, plainText)
	}
}

func TestDecryptInvalidKey(t *testing.T) {
	plainText := "HELP"
	key := "DDCF"
	cipherText, err := Encrypt(plainText, key)
	if err != nil {
		t.Errorf("Encrypt returned error: %v", err)
	}

	_, err = Decrypt(cipherText, "InvalidKey")
	if err == nil {
		t.Errorf("Decrypt did not return error for invalid key")
	}
}


