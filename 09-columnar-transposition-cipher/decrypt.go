package main

import (
  //,"strings"
  "fmt"
  "math"
  "errors"
)

func decrypt (ciphertext, key string) (string, error) {

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

/* 
  indexMap := make(map[string]int)
  for i, val := range key {
    indexMap[val] = i
  }
*/

  index := 0
  for i, _ := range sortedKey {
    for row := 0; row < numRows; row++ {
      if index < len(ciphertext) {
        table[row][i] = rune(ciphertext[index])
        index++
      } else {
        break
      }
    }
  }

  fmt.Println(table)

  var plaintext string
	for _, row := range table {
		for _, r := range row {
			if r != 0 {
				plaintext += string(r)
			}
		}
	}
	return plaintext, nil
}
