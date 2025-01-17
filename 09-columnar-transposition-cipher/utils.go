package main

import (
  "strings"
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
