
# XOR Cipher Implementation

The XOR cipher is a simple, yet effective encryption technique that has been used for centuries. Its origins date back to the early days of cryptography, when it was used by ancient civilizations such as the Egyptians and Greeks.

## Background

The XOR cipher works by performing a bitwise exclusive or (XOR) operation between the plaintext message and a secret key. This operation produces a ciphertext that can only be decrypted by performing the same XOR operation with the same key.

## Implementation

This implementation provides two versions of the XOR cipher:

1. *Binary XOR Cipher*: This implementation uses a binary representation of the plaintext message and key. It performs a bitwise XOR operation between the two binary strings to produce the ciphertext.
2. *Gob XOR Cipher*: This implementation uses the Go `encoding/gob` package to encode the plaintext message and key into binary format. It then performs a bitwise XOR operation between the two binary strings to produce the ciphertext.

## Usage

To use the XOR cipher, simply call the `Encrypt` or `xorCipher` function with the plaintext message and key as arguments. The resulting ciphertext can be decrypted using the `Decrypt` or `xorDecipher` function with the same key.

*Example*

```
package main

import (
	"fmt"
	"xorcipher"
)

func main() {
	plaintext := "Hello, World!"
	key := "secretkey"

	ciphertext, err := xorcipher.Encrypt(plaintext, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Ciphertext:", ciphertext)

	decrypted, err := xorcipher.Decrypt(ciphertext, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Decrypted:", decrypted)
}
```

## Security

The XOR cipher is a simple and efficient encryption technique, but it is not considered secure for protecting sensitive information. It is vulnerable to frequency analysis attacks and can be easily broken by an attacker with access to the ciphertext and a sample of the plaintext.


