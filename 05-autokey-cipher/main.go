package main

import (
  "strings"
  "fmt"
  "regexp"
  "errors"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main () {
  enc, err := Encipher("HELLO", "N")
  if err == nil {
    fmt.Println(enc)
  }
  enc, err = Encipher("GeEkSfOrGeEkS", "P")
  if err == nil {
    fmt.Println(enc)
  }

  dec, err := Decipher("uLpWz", "n")
  if err == nil {
    fmt.Println(dec)
  }
}

func Encipher (plainText string, key string) (string, error) {
  if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(plainText) {
    return "", errors.New("Input string must be a non-empty string containing only letters and/or spaces")
  }

  if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(key) {
   return "", errors.New("key can only contain alphabet characters")
  }
  plainText, key = strings.ToUpper(strings.ReplaceAll(plainText, " ", "")), strings.ToUpper(fillShortString(plainText, key))

  encryptedText := ""
  for i, c := range plainText {

    encryptedText += string(alphabet[(int(c) - 65 + int(key[i % len(key)]) - 65) % 26])
  }

  return encryptedText, nil
}

func fillShortString(longStr, shortStr string) string {
	longStrLen := len(longStr)
	shortStrLen := len(shortStr)

	if longStrLen <= shortStrLen {
		return shortStr[:longStrLen]
	}

	fillLen := longStrLen - shortStrLen
	fillStr := longStr[:fillLen]

	return shortStr + fillStr
}

func Decipher (cipherText, key string) (string, error) {
  if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(cipherText) {
    return "", errors.New("Input string must be a non-empty string containing only letters and/or spaces")
  }

  if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(key) {
   return "", errors.New("key can only contain alphabet characters")
  }
  cipherText, key = strings.ToUpper(strings.ReplaceAll(cipherText, " ", "")), strings.ToUpper(key)

  fmt.Println(cipherText, key)
  decryptedText := ""
  for i, c := range cipherText {

    decryptedText += string(alphabet[(int(c) - int(key[i % len(key)]) + 26) % 26])
    key += string(decryptedText[i])
  }

  return decryptedText, nil
}
