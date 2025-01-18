package main

import (
  "errors"
  "math"
)

func encrypt (text, key string) (string, error) {
  
  if len(text) == 0 || len(key) == 0 {
    return "", errors.New("Text and key must contain at least one character!")
  }

  errText, errKey := checkAlphabet(text), checkAlphabet(key)
  if errText != nil {
    return "", errText
  } 
  if errKey != nil {
    return "", errKey
  }

  text, key = removeSpaces(toUpperCase(text)), removeSpaces(toUpperCase(key))
  numColumns := len(key)
  numRows := int(math.Ceil(float64(len(text)) / float64(numColumns)))

  table := make([][]rune, numRows)
  for i := range table {
    table[i] = make([]rune, numColumns)
  }

  // filling the transposition table …
  index := 0
  for i := 0; i < numRows; i++ {
    for j := 0; j < numColumns; j++ {
      if index < len(text) {
        table[i][j] = rune(text[index])
        index++
      } else {
        table[i][j] = 'X'
      }
    }
  }

  perm := sortedKeyIndices(key)

	var encryptedText string
	for _, colIndex := range perm {
		for _, row := range table {
			if colIndex < len(row) {
				encryptedText += string(row[colIndex])
			}
		}
	}

	return encryptedText, nil
}
