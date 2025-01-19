package main

import (
  "strings"
  "errors"
  "regexp"
  "sort"
)

func sortWord(word string) string {
	runeSlice := []rune(word)
	for i := 0; i < len(runeSlice); i++ {
		for j := i + 1; j < len(runeSlice); j++ {
			if runeSlice[i] > runeSlice[j] {
				runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
			}
		}
	}
	return string(runeSlice)
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

func removeSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' {
			return -1
		}
		return r
	}, s)
}

type indexedChar struct {
    index int
    char  rune
}

func sortedKeyIndices(key string) []int {
    indexedChars := make([]indexedChar, len(key))
    for i, char := range key {
        indexedChars[i] = indexedChar{i, char}
    }
    /*
    sort.Slice(indexedChars, func(i, j int) bool {
        return indexedChars[i].char < indexedChars[j].char
    }) */

    sort.Slice(indexedChars, func(i, j int) bool {
        if indexedChars[i].char == indexedChars[j].char {
            return indexedChars[i].index < indexedChars[j].index
        }
        return indexedChars[i].char < indexedChars[j].char
    })


    var sortedIndices []int
    for _, indexedChar := range indexedChars {
        sortedIndices = append(sortedIndices, indexedChar.index)
    }
    return sortedIndices
}

func checkAlphabet(s string) error {
	match, err := regexp.MatchString("^[a-zA-Z\\s]+$", s)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("string contains non-alphabet characters")
	}
	return nil
}
