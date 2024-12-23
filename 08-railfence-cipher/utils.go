package main

import (
	"math"
)

func FormRails(text string, rails int) [][]rune {
	//start := 0
	lines := make([][]rune, rails)

	for i := range lines {
		lines[i] = make([]rune, 0, int(math.Ceil(float64(len(text)/rails))))
	}

	/*
		for k := 0; k < rails; k++ {
			for i := k; i < len(text); i += rails {
				lines[k] = append(lines[k], rune(text[i]))
			}
		}
	*/

	dir := 1 // direction: 1 for down, -1 for up
	rail := 0
	for i := 0; i < len(text); i++ {
		lines[rail] = append(lines[rail], rune(text[i]))
		if rail == 0 {
			dir = 1
		} else if rail == rails-1 {
			dir = -1
		}
		rail += dir
	}

	return lines
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

func removeSpaces (s string) string {
	var result string
	for _, char := range s {
		if char != ' ' {
			result += string(char)
		}
	}
	return result
}
