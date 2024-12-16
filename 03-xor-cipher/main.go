package main

import (
    "fmt"
    "strconv"
    "errors"
)

func main () {
}

func binary(s string) string {
    for _, c := range s {
        if c != '0' && c != '1' {
            res := ""
            for _, c := range s {
                res = fmt.Sprintf("%s%.8b", res, c)
            }
            return res
        }
    }
    return s
}
/** 
    `Encrypt` encrypts a plaintext message using a plaintext key. The functiom uses a simple XOR
    Cipher.
*/
func Encrypt (plainText string, key string) (string, error) {
    if (len(key)) == 0 {
        return "", errors.New("Key cannot be empty!")
    }
    var textBin, keyBin string = binary(plainText), binary(key)
    cipherText := ""

    for i, _ := range textBin {
        if textBin[i] == keyBin[i % len(keyBin)] {
            cipherText += "0"
        } else {
            cipherText += "1"
        }
    }

    return cipherText, nil
}
/**
`Decrypt` decrypts a cipher text using the plaintext key. The function uses a simple XOR Cipher. It
returns an error if the ciphertext is not a valid binary string.
*/
func Decrypt (cipherText string, key string) (string, error) {
	for _, c := range cipherText {
		if c != '0' && c != '1' {
			return "", errors.New("cipher text is not a valid binary string")
		}
	}
	return Encrypt(cipherText, key)
}

func binaryToString(binary string) string {
	var result string
	for len(binary) > 0 {
		if len(binary) < 8 {
			binary = padBinary(binary)
		}
		byteStr := binary[:8]
		binary = binary[8:]
		byteValue, _ := strconv.ParseInt(byteStr, 2, 8)
		result += string(byte(byteValue))
	}
	return result
}

func padBinary(binary string) string {
	for len(binary) < 8 {
		binary = "0" + binary
	}
	return binary
}
