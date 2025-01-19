package main

import "errors"

// Decrypt decrypts the given ciphertext using the rail fence cipher with the specified number of rails.
func Decrypt (ciphertext string, rails int) (string, error) {
  
  if len(ciphertext) == 0 || rails <= 0 {
    return "", errors.New("Invalid ciphertext, and/or rail length provided.")
  }

  if rails <= 1 {
		return ciphertext, nil
  }

  n := len(ciphertext)
	pattern := make([]int, n)
	row, direction := 0, 1

  for i := 0; i < n; i++ {
		pattern[i] = row
		if row == 0 {
			direction = 1
		} else if row == rails-1 {
			direction = -1
		}
		row += direction
  }

  railLengths := make([]int, rails)
	for _, r := range pattern {
		railLengths[r]++
	}

	railsContent := make([][]rune, rails)
	index := 0
	for r := 0; r < rails; r++ {
		railsContent[r] = []rune(ciphertext[index : index+railLengths[r]])
		index += railLengths[r]
  }

  plaintext := make([]rune, n)
	rowPointers := make([]int, rails)
	row, direction = 0, 1

	for i := 0; i < n; i++ {
		plaintext[i] = railsContent[row][rowPointers[row]]
		rowPointers[row]++
		if row == 0 {
			direction = 1
		} else if row == rails-1 {
			direction = -1
		}
		row += direction
	}

	return string(plaintext), nil
}
