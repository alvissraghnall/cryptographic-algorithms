# Autokey Cipher Implementation

This is a Go implementation of the Autokey cipher, a polyalphabetic substitution cipher that uses a keyword and parts of the plaintext to encrypt and decrypt messages.

### _Implementation Details

This implementation of the Autokey cipher uses a keyword to start the encryption process, but then uses the plaintext itself to continue the encryption. The cipher eliminates spaces from the input string before encrypting or decrypting it.

### Usage

To use this implementation, simply call the `Encipher` or `Decipher` functions, passing in the input string and keyword as arguments. For example:
```
enc, err := Encipher("HELLO", "N")
if err == nil {
    fmt.Println(enc)
}
```
### Functions

- `Encipher(plainText string, key string) (string, error)`: Encrypts the input string using the Autokey cipher.
- `Decipher(cipherText string, key string) (string, error)`: Decrypts the input string using the Autokey cipher.
- `fillShortString(longStr string, shortStr string) string`: Fills the short string with characters from the long string to match its length.

### Notes

- The keyword, and plaintext should only contain letters.
- This implementation eliminates spaces from the input string before encrypting or decrypting it.
