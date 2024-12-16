package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(rune('T'))
	Encrypt("THISISWIKIPEDIA", "CIPHER")
}

func Encrypt(plainText string, key string) (string, error) {
	if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(plainText) {
		return "", errors.New("Input string must be a non-empty string containing only letters and/or spaces")
	}

	columnsNum := int(math.Ceil(float64(len(plainText)) / float64(len(key))))
	holder := make([][]rune, columnsNum)

	fmt.Printf("%v %v %v\n", columnsNum, holder, key)

	/**
	  for i, _ := range key {
	    holder[i] = []rune(plainText[i*len(key):i*len(key)+len(key)])
	  }
	*/

	for i := range columnsNum {
		start := i * len(key)
		end := start + len(key)
		if end > len(plainText) {
			remaining := end - len(plainText)
			holder[i] = []rune(plainText[start:len(plainText)] + strings.Repeat("_", remaining))
		} else {
			holder[i] = []rune(plainText[start:end])
		}
		fmt.Printf("%x %v\n", string(holder[i][0]), start)
	}

	fmt.Printf("%v\n", holder)

	type CharInfo struct {
		char  rune
		index int
	}

	charInfos := make([]CharInfo, 0, len(key))

	for i, c := range key {
		charInfos = append(charInfos, CharInfo{char: c, index: i})

		sort.Slice(charInfos, func(i, j int) bool {
			return charInfos[i].char <= charInfos[j].char
		})
	}

	var sortedKey string
	for _, ci := range charInfos {
		sortedKey += string(ci.char)
	}

	fmt.Println(sortedKey)

	var cipherText = make([][]rune, len(key))
	for i := range cipherText {
		cipherText[i] = make([]rune, columnsNum)
	}

	charInfoMap := make(map[rune]CharInfo)
	for _, ci := range charInfos {
		charInfoMap[ci.char] = ci
	}

	for idx, letter := range sortedKey {
		ci, ok := charInfoMap[letter]
		if ok {
			index := ci.index
			for i, outer := range holder {
				value := outer[index]
				cipherText[idx][i] = value
			}
		}
	}

	fmt.Printf("%v %v\n", cipherText, string(cipherText[0]))

	return "", nil
}
