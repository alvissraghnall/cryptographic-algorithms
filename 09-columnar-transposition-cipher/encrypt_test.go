package main

import (
	"testing"
	"errors"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		key       string
		expected  string
		expectErr bool
		errMsg    string
	}{
		// Normal cases
		{"WE ARE DISCOVERED FLEE AT ONCE", "ZEBRA", "EODAEASRENEIELORCEECWDVFT", false, ""},
		{"ATTACK AT DAWN", "KEY", "TCTWAAAATKDN", false, ""}, // Testing with key shorter than plaintext

		// Edge cases
		{"HELLO", "WORLDS", "OLELXH", false, ""},             // Plaintext shorter than key
		{"COLUMNAR TRANSPOSITION", "NOITISOPSNARTRANMULOC", "AONLMIICRSOAORNPNTUST", false, "" }, // Key is same length as plaintext
		{"", "EMPTY", "", true, "Text and key must contain at least one character!" },                           // Empty plaintext
		{"EMPTYTEXT", "", "EMPTYTEXT", true, "Text and key must contain at least one character!" },             // empty key

		// Plaintext and key with mixed cases
		{"We Are Hidden", "KeY", "EEDNWRIEAHDX", false, ""}, // Mixed case plaintext and key

		// Large plaintext
		{
			"THIS IS A LONGER TEXT THAT WE ARE ENCRYPTING USING THE COLUMNAR TRANSPOSITION CIPHER",
			"CIPHER",
			"TARHRYUHMAIIIGTECNGLTONRSNXWNINORPOEHLTAEPSENNTPIOETETICASIHSETARGTURSCX",
			false,
			"",
		},

		// Plaintext with spaces and special characters
		{"THIS IS A TEST! 1234", "SIMPLE", "IIAHSSTEXTL1SXX23T4", true, "string contains non-alphabet characters" },

		// Key containing repeated characters
		{"WE ARE DISCOVERED FLEE AT ONCE", "ZZEBRA", "DEECXROFOXACDTXEVLNXWIREEESEAX", false, ""},
	}

	for _, test := range tests {
		actual, err := encrypt(test.plaintext, test.key)
		if test.expectErr {
			// Check if error occurred as expected
			if err == nil {
				t.Errorf("encrypt(%q, %q) expected error, got none", test.plaintext, test.key)
			} else if !errors.Is(err, errors.New(test.errMsg)) {
				t.Errorf("encrypt(%q, %q) expected error %q, got %q", test.plaintext, test.key, test.errMsg, err.Error())
			}
		} else {
			// Ensure no error occurred for valid inputs
			if err != nil {
				t.Errorf("encrypt(%q, %q) unexpected error: %v", test.plaintext, test.key, err)
			}
			// Check if the output matches the expected value
			if actual != test.expected {
				t.Errorf("encrypt(%q, %q) = %q; want %q", test.plaintext, test.key, actual, test.expected)
			}
		}
	}
}
