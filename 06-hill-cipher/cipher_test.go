
package main

import (
	"testing"
)

func TestGetKeyMatrix(t *testing.T) {
	key := "ABCDEFGHI"
	matrix, err := GetKeyMatrix(key)
	if err != nil {
		t.Errorf("GetKeyMatrix returned error: %v", err)
	}
	if len(matrix) != 3 || len(matrix[0]) != 3 {
		t.Errorf("GetKeyMatrix returned matrix with incorrect dimensions: %v", matrix)
	}
}

func TestGetKeyMatrixInvalidKey(t *testing.T) {
	key := "ABCDEF"
	_, err := GetKeyMatrix(key)
	if err == nil {
		t.Errorf("GetKeyMatrix did not return error for invalid key")
	}
}

func TestParseTextToVector(t *testing.T) {
	plainText := "ACT"
	vector, err := ParseTextToVector(plainText, 3)
	if err != nil {
		t.Errorf("ParseTextToVector returned error: %v", err)
	}
	if len(vector) != 1 || len(vector[0]) != 3 {
		t.Errorf("ParseTextToVector returned vector with incorrect dimensions: %v", vector)
	}
}

func TestParseTextToVectorEmptyText(t *testing.T) {
	plainText := ""
	_, err := ParseTextToVector(plainText, 3)
	if err == nil {
		t.Errorf("ParseTextToVector did not return error for empty text")
	}
}

func TestEncrypt(t *testing.T) {
	plainText := "HELP"
	key := "DDCF"
	cipherText, err := Encrypt(plainText, key)
	if err != nil {
		t.Errorf("Encrypt returned error: %v", err)
	}
	if len(cipherText) != 4 {
		t.Errorf("Encrypt returned cipher text with incorrect length: %v", cipherText)
	}
}

func TestEncryptInvalidKey(t *testing.T) {
	plainText := "HELP"
	key := "DDC"
	_, err := Encrypt(plainText, key)
	if err == nil {
		t.Errorf("Encrypt did not return error for invalid key")
	}
}

func TestMatrixVectorMultiplication(t *testing.T) {
	matrix := Matrix{
		{1, 2},
		{3, 4},
	}
	vector := Vector{5, 6}
	result, err := MatrixVectorMultiplication(matrix, vector)
	if err != nil {
		t.Errorf("MatrixVectorMultiplication returned error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("MatrixVectorMultiplication returned result with incorrect length: %v", result)
	}
}

func TestMatrixVectorMultiplicationInvalidDimensions(t *testing.T) {
	matrix := Matrix{
		{1, 2},
		{3, 4},
	}
	vector := Vector{5}
	_, err := MatrixVectorMultiplication(matrix, vector)
	if err == nil {
		t.Errorf("MatrixVectorMultiplication did not return error for invalid dimensions")
	}
}
