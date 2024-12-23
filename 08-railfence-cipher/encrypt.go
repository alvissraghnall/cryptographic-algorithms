
package main

import (
  "errors"
)

func Encrypt (text string, rails int) (string, error) {
  if len(text) == 0 || rails <= 0 {
    return "", errors.New("Invalid plaintext, and/or rail length provided.")
  }

  text = removeSpaces(toUpperCase(text))

  cipher := FormRails(text, rails)

  cipherText := ""
    for i := range cipher {
      cipherText += string(cipher[i])
    }

  return cipherText, nil
}
