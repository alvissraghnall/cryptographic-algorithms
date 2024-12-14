

# Vigenère Cipher Implementation in Go

This repository contains a Go implementation of the Vigenère Cipher, a polyalphabetic substitution cipher that uses a keyword to encrypt and decrypt messages.

## About the Vigenère Cipher

The Vigenère Cipher is named after Giovan Battista Bellaso's associate, Blaise de Vigenère, who described it in his book "Traicté des Chiffres" in 1586. However, the cipher was actually first described by Bellaso in his book "La Cifra" in 1550.

The Vigenère Cipher was considered unbreakable for many years, earning it the nickname "undecipherable cipher." However, in 1863, Friedrich Kasiski published a method for breaking the cipher, and in 1917, William Friedman developed a more efficient method.

## Features

- Encrypts and decrypts messages using the Vigenère Cipher algorithm
- Supports repeating the keyword to match the length of the message
- Includes error handling for invalid characters in the message and keyword

## Usage 

To use the Vigenère Cipher implementation, simply call the `Encrypt` or `Decrypt` function with the message and keyword as arguments.

```
message := "ATTACKATDAWN"
keyword := "LEMON"

encryptedMessage, err := Encrypt(message, keyword)
if err != nil {
    fmt.Println(err)
    return
}

decryptedMessage, err := Decrypt(encryptedMessage, keyword)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println("Encrypted Message:", encryptedMessage)
fmt.Println("Decrypted Message:", decryptedMessage)
```

