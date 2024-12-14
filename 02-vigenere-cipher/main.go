
package main

import (
  "strings"
  "fmt"
  "errors"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func repeatKey (key string, messageLen int) (string, error) {
  var newKey string = ""
  if !IsLetter(key) {
    return "", errors.New("Key contains invalid characters.")
  }

  for len(newKey) < messageLen {
    newKey += key
  }

  return newKey[:len(newKey)-len(newKey) % messageLen], nil
}


func IsLetter(s string) bool {
  return !strings.ContainsFunc(s, func(r rune) bool {
    return r < 'A' || r > 'Z'
  })
}

func checkEmptyString (s string) bool {
  return len(s) == 0
}

func Encrypt (text string, key string) (string, error) {

  if (checkEmptyString(text) || checkEmptyString(key)) {
    return "", errors.New("Cannot encrypt empty string, or use empty string as key")
  }

  if !IsLetter (text) || !IsLetter(key) {
    return "", errors.New("text/key contains invalid characters")
  }


  var encryptedText string = ""

  for i, c := range text {

    encryptedText += string(alphabet[(int(c) - 65 + int(key[i % len(key)]) - 65) % 26])
  }

  return encryptedText, nil
}

func Decrypt (cipherText string, key string) (string, error) {
  if (checkEmptyString(cipherText) || checkEmptyString(key)) {
    return "", errors.New("Cannot encrypt empty string, or use empty string as key")
  }

  if !IsLetter (cipherText) || !IsLetter(key) {
    return "", errors.New("text/key contains invalid characters")
  }

  var decryptedText string = ""

  for i, c := range cipherText {
    //fmt.Printf("%v %v\n", int(c), int(key[i % len(key)]))
    //fmt.Println((int(c) - int(key[i % len(key)]) + 26) % 26)
    decryptedText += string(alphabet[(int(c) - int(key[i % len(key)]) + 26) % 26 ])
  }

  return decryptedText, nil
}


func main () {

    message := "ATTACKATDAWN"
    keyword := "LEMON"
    encryptedMessage, err := Encrypt(message, keyword)
    if err != nil {
        fmt.Printf("Test case 1 failed: %v\n", err)
        return
    }
    decryptedMessage, err := Decrypt(encryptedMessage, keyword)
    if err != nil {
        fmt.Printf("Test case 1 failed: %v\n", err)
        return
    }
    if decryptedMessage != message {
        fmt.Println("Test case 1 failed")
        return
    }
    fmt.Println("Test case 1 passed")
 
    // Test case 2: Encryption and decryption with repeating keyword
    message = "THISISASECRETMESSAGE"
    keyword = "CODE"
    encryptedMessage, err = Encrypt(message, keyword)
    if err != nil {
        fmt.Printf("Test case 2 failed: %v\n", err)
        return
    }
    decryptedMessage, err = Decrypt(encryptedMessage, keyword)
    if err != nil {
        fmt.Printf("Test case 2 failed: %v\n", err)
        return
    }
    if decryptedMessage != message {
        fmt.Println("Test case 2 failed")
        return
    }
    fmt.Println("Test case 2 passed")

    // Test case 3: Encryption and decryption with single-character keyword
    message = "HELLO"
    keyword = "A"
    encryptedMessage, err = Encrypt(message, keyword)
    if err != nil {
        fmt.Printf("Test case 3 failed: %v\n", err)
        return
    }
    decryptedMessage, err = Decrypt(encryptedMessage, keyword)
    if err != nil {
        fmt.Printf("Test case 3 failed: %v\n", err)
        return
    }
    if decryptedMessage != message {
        fmt.Println("Test case 3 failed")
        return
    }
    fmt.Println("Test case 3 passed")

    // Test case 4: Encryption and decryption with single-character message
    message = "A"
    keyword = "CODE"
    encryptedMessage, err = Encrypt(message, keyword)
    if err != nil {
        fmt.Printf("Test case 4 failed: %v\n", err)
        return
    }
    decryptedMessage, err = Decrypt(encryptedMessage, keyword)
    if err != nil {
        fmt.Printf("Test case 4 failed: %v\n", err)
        return
    }
    if decryptedMessage != message {
        fmt.Println("Test case 4 failed")
        return
    }
    fmt.Println("Test case 4 passed")

    // Test case 5: Encryption and decryption with empty message
    message = ""
    keyword = "CODE"
    
    _, err = Encrypt(message, keyword)
    if err.Error() != "Cannot encrypt empty string, or use empty string as key" {
        fmt.Println("Test case 6 failed")
        return
    }
    fmt.Println("Test case 5 passed")

    // Test case 6: Encryption and decryption with invalid characters
    message = "HELLO!"
    keyword = "CODE"
    _, err = Encrypt(message, keyword)
    if err.Error() != "text/key contains invalid characters" {
        fmt.Println("Test case 6 failed")
        return
    }
    fmt.Println("Test case 6 passed")


    // Test case 7: Encryption and decryption with keyword containing invalid characters
    message = "HELLO"
    keyword = "CODE!"
    _, err = Encrypt(message, keyword)
    if err.Error() != "text/key contains invalid characters" {
        fmt.Println("Test case 7 failed")
        return
    }
    fmt.Println("Test case 7 passed")

    // Test case 8: Encryption and decryption with empty keyword
    message = "HELLO"
    keyword = ""
    _, err = Encrypt(message, keyword)
    if err.Error() != "Cannot encrypt empty string, or use empty string as key" {
        fmt.Println("Test case 8 failed")
        return
    }
    fmt.Println("Test case 8 passed")

    fmt.Println("All test cases passed")
}
