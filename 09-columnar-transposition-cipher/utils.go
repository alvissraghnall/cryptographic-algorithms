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

func sortedKeyIndices(key string) []int {
	sortedIndices := make([]int, len(key))
	for i := range sortedIndices {
		sortedIndices[i] = i
	}
	sort.Slice(sortedIndices, func(i, j int) bool {
		return key[sortedIndices[i]] < key[sortedIndices[j]]
	})
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
