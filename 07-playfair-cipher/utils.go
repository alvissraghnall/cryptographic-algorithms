package main

import (
	"errors"
	"strings"
	"fmt"
)

type Grid [][]rune
const allLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func createKeyGrid (key string) (Grid, error) {

	if len(key) == 0 || len(key) > 25 {
		return nil, errors.New("Key cannot be empty or consist of over 25 characters!")
	}

	uniqueLettersFromKey, _ := uniqueLetters(key)

	keyGrid := make(Grid, 5)
	
	for i := range keyGrid {
		keyGrid[i] = make([]rune, 5)
	}

	complete := completeAlphabet(uniqueLettersFromKey)

	keyGrid.FormRows(complete)
	return keyGrid, nil
}

func uniqueLetters(s string) (string, map[string]bool) {
    unique := make(map[string]bool)
    order := ""
    for _, c := range s {
        letter := string(c)
        if !unique[letter] {
            unique[letter] = true
            order += letter
        }
    }
    return order, unique
}

func completeAlphabet(unique string) string {
	for _, c := range allLetters {
		if !strings.Contains(unique, string(c)) && c != 'J' {
			unique += string(c)
		}
	}
	return unique
}

func (grid *Grid) FormRows(str string) {
	fmt.Println(str, rune(str[0]))
    k := 0
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
						//fmt.Println(rune(str[k]))
            if k < len(str) {
                (*grid)[i][j] = rune(str[k])
                k++
            } else {
                break
            }
        }
    }
}

func toUpperCase(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 96 && rune(s[i]) < 123 {
			value := rune(s[i]) - 32
			result += string(value)
		} else {
			result += string(s[i])
		}
	}
	return result
}

func splitIntoDigraphs(plaintext string) [][]string {
	// Convert plaintext to uppercase
	plaintext = toUpperCase(plaintext)

	// Split plaintext into individual letters
	letters := strings.Split(plaintext, "")

	// If there's an odd number of letters, add Z
	if len(letters)%2 != 0 {
		letters = append(letters, "Z")
	}

	// Initialize digraphs slice
	digraphs := make([][]string, 0)

	// Iterate over letters in steps of 2
	for i := 0; i < len(letters); i += 2 {
		// If the current letter is the same as the next one, insert X
		if letters[i] == letters[i+1] {
			letters = append(letters[:i+1], append([]string{"X"}, letters[i+1:]...)...)
			digraphs = append(digraphs, []string{letters[i], "X"})
			i++
		} else {
			digraphs = append(digraphs, []string{letters[i], letters[i+1]})
		}
	}

	return digraphs
}

