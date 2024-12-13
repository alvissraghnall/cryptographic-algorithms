package main

 import (
  "fmt"
  "strings"
  "errors"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var textPlusShift int

type ShiftDirection int

const (
  LEFT ShiftDirection = iota
  RIGHT
)

func encrypt (text string, shift int, shiftDirection ShiftDirection) (string, error) {

  if shift < 0 {
    return "", errors.New("shift value cannot be negative")
  }

  if shiftDirection != LEFT && shiftDirection != RIGHT {
    return "", errors.New("invalid shift direction")
  }

  if !strings.ContainsAny(text, alphabet+" ") {
    return "", errors.New("text contains invalid characters")
  }


  var encryptedText string = ""
  
  for _, c := range text {
    if (c == ' ') { 
      encryptedText += " "
      continue 
    }

    /** En(x) = (x + n) % 26
  
      The rune type in Go is basically an alias for int32, and it holds individual characters
      in a program. Since the rune type is basically ints, every utf8 character (i think)
      has an integer representation. For 'A', it is 65. 'B' is 66, and so on...

      So, here, we take the current character from the plaintext to be encrypted and parse
      it to an int (Go is sorta strict with types lol), and then subtract 65 from it (Since
      the integer values for runes relevant to us — uppercase characters start at 65, so 
      we can index our alphabet array). the rest is pretty straightforward. 
    */

    if(shiftDirection == LEFT) {
      textPlusShift = int(c) - 65 - shift + 26
    } else {
      textPlusShift = int(c) - 65 + shift
    }
    
    encryptedText += string(alphabet[textPlusShift % 26])
  }

  return encryptedText, nil
}

func decrypt (text string, shift int, shiftDirection ShiftDirection) (string, error) {
  decryptedText := ""
  
  for _, c := range text {
    if (c == ' ') {
      decryptedText += " "
      continue
    }

    if(shiftDirection == LEFT) {
      textPlusShift = int(c) - 65 + shift
    } else {
      textPlusShift = int(c) - 65 - shift + 26
    }
  
    decryptedText += string(alphabet[textPlusShift % 26])
   
  }

  return decryptedText, nil
}

func main () {
   text := "HELLO"
        shift := 3
        shiftDirection := RIGHT

        encryptedText, err := encrypt(text, shift, shiftDirection)
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Println("Encrypted Text:", encryptedText)

        decryptedText, err := decrypt(encryptedText, shift, shiftDirection)
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Println("Decrypted Text:", decryptedText)
}
