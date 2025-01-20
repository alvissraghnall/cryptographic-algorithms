package main

import (
  "strings"
  "errors"
)

func encryptPlayfair(plaintext, key string) (string, error) {

    key = replaceChars(toUpperCase(key)) // Treat 'J' as 'I' in the Playfair cipher & remove spaces
    keyGrid, err := createKeyGrid(key)
    if err != nil {
        return "", err
    }

    plaintext = toUpperCase(plaintext)
    plaintext = replaceChars(plaintext)
    digraphs := splitIntoDigraphs(plaintext)

    var ciphertext strings.Builder

    // Encrypt each digraph
    for _, digraph := range digraphs {
        if len(digraph) != 2 {
            continue
        }
        char1, char2 := rune(digraph[0][0]), rune(digraph[1][0])
        r1, r2 := findPosition(keyGrid, char1), findPosition(keyGrid, char2)

        if r1 == nil || r2 == nil {
            return "", errors.New("invalid character in plaintext")
        }

        switch {
            case r1[0] == r2[0]: // Same row
                ciphertext.WriteRune(keyGrid[r1[0]][(r1[1]+1)%5])
                ciphertext.WriteRune(keyGrid[r2[0]][(r2[1]+1)%5])
            case r1[1] == r2[1]: // Same column
                ciphertext.WriteRune(keyGrid[(r1[0]+1)%5][r1[1]])
                ciphertext.WriteRune(keyGrid[(r2[0]+1)%5][r2[1]])
            default: // Rectangle swap
                ciphertext.WriteRune(keyGrid[r1[0]][r2[1]])
                ciphertext.WriteRune(keyGrid[r2[0]][r1[1]])
        }
    }

        return ciphertext.String(), nil
}

// Helper function to find the position of a character in the key grid
func findPosition(grid Grid, char rune) (position []int) {
    for i, row := range grid {
        for j, cell := range row {
            if cell == char {
                return []int{i, j}
            }
        }
    }
    return nil
}

