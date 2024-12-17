package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)


type CharInfo struct {
	char  rune
	index int
}

func main() {
	fmt.Println(rune('T'))
	encrypted, err := Encrypt("THISISWIKIPEDIA", "CIPHER")

	if err == nil {
		fmt.Println(encrypted)
	}

	decrypted, err := Decrypt("TWDIP_SI_HIIIKASE_", "CIPHER")
	if err == nil {
		fmt.Println(decrypted)
	}
}

func Encrypt(plainText string, key string) (string, error) {
	if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(plainText) {
		return "", errors.New("Input string must be a non-empty string containing only letters and/or spaces")
	}

  if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(key) {
		return "", errors.New("key can only contain alphabet characters")
  }

	plainText, key = strings.ToUpper(plainText), strings.ToUpper(key)

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

	var result string
  for _, row := range cipherText {
		result += string(row)
	}

	fmt.Printf("%v %v\n", cipherText, string(cipherText[0]))

	return result, nil
}

func Decrypt(cipherText string, key string) (string, error) {
    if !regexp.MustCompile(`^[a-zA-Z_]+$`).MatchString(cipherText) {
        return "", errors.New("cipherText can only contain alphabet characters and underscores")
    }
    if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(key) {
        return "", errors.New("Key can only contain alphabet characters")
    }

		numOfRows := int(math.Ceil(float64(len(cipherText) / len(key))))
    
		holder := make([][]rune, numOfRows)

		for i := range numOfRows {
			start := i * len(key)
			end := start + len(key)
			if end > len(cipherText) {
				remaining := end - len(cipherText)
				holder[i] = []rune(cipherText[start:len(cipherText)] + strings.Repeat("_", remaining))
			} else {
				holder[i] = []rune(cipherText[start:end])
			}
		}	
		fmt.Printf("%v\n", holder)

		charInfos := make([]CharInfo, len(key))
for i, c := range key {
    charInfos[i] = CharInfo{char: c, index: i}
}
//sort.Slice(charInfos, func(i, j int) bool { return charInfos[i].char <= charInfos[j].char })

decipheredText := make([][]rune, len(key))
for i := range decipheredText {
    decipheredText[i] = make([]rune, numOfRows)
}

charInfoMap := make(map[rune]CharInfo)
        for _, ci := range charInfos {
                charInfoMap[ci.char] = ci
        }

        for idx, letter := range key {
                ci, ok := charInfoMap[letter]
                if ok {
                        index := ci.index
                        for i, outer := range holder {
                                value := outer[index]
                                decipheredText[idx][i] = value
                        }
                }
        }


var result string
for _, row := range decipheredText { result += string(row)
}
return result, nil

}
