package main

import (
  "strings"
  "fmt"
  "math"
)

func decipher (ciphertext, key string) (string, error) {

  if len(ciphertext) == 0 || len(key) == 0 {
    return "", errors.New("Text and key must contain at least one character!")
  }

  errText, errKey := checkAlphabet(ciphertext), checkAlphabet(key)
  if errText != nil {
    return "", errText
  }
  if errKey != nil {
    return "", errKey
  }

  ciphertext, key = removeSpaces(toUpperCase(ciphertext)), removeSpaces(toUpperCase(key))
  numColumns := len(key)
  numRows := int(math.Ceil(float64(len(ciphertext)) / float64(numColumns)))

  sortedKey := sortWord(key)

  table := make([][]rune, numRows)
  for i := range table {
    table[i] = make([]rune, numColumns)
  }
  
  index := 0
  for _, col := range sortedKey {
    for row := 0; row < numRows; row++ {
      if index < len(ciphertext) {
        table[row][strings.Index(sortedKey, string(col))] = rune(ciphertext[index])
        index++
      } else {
        break
      }
    }
  }

  fmt.Println(table)

  return "", nil
}
