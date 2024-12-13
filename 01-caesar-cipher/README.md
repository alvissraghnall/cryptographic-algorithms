
#Caesar Cipher Implementation in Go

This folder contains a simple implementation of the Caesar Cipher algorithm in Go. 
The Caesar Cipher is a type of substitution cipher where each letter in the plaintext is 'shifted' a certain number of places down the alphabet.

*Features*

- Encrypts and decrypts text using the Caesar Cipher algorithm
- Supports shifting in both left and right directions
- Handles uppercase letters and spaces

*Usage*

To use this implementation, simply call the `encrypt` or `decrypt` function with the text to be encrypted or decrypted, the shift value, and the shift direction.

```
encryptedText := encrypt("HELLO", 3, RIGHT)
decryptedText := decrypt(encryptedText, 3, RIGHT)
```

*Code Structure*

The code is organized into two main functions:

- `encrypt`: Encrypts the input text using the Caesar Cipher algorithm
- `decrypt`: Decrypts the input text using the Caesar Cipher algorithm

*Notes*

- This implementation only handles uppercase letters and spaces. If you need to handle lowercase letters or other characters, you will need to modify the code accordingly.
