# Hill Cipher Implementation

### Table of Contents

1. Introduction
2. What is the Hill Cipher?
3. How does it work?
4. Implementation Details
5. Usage
6. Testing
7. Lore

**Introduction**

This is an implementation of the Hill Cipher, a polygraphic substitution cipher that uses linear algebra to encrypt and decrypt messages.

**What is the Hill Cipher?**

The Hill Cipher is a polygraphic substitution cipher that was invented by Lester S. Hill in 1929. It is a block cipher that operates on fixed-length blocks of plaintext and ciphertext.

**How does it work?**

The Hill Cipher works by converting each block of plaintext into a vector, and then multiplying that vector by a key matrix to produce the ciphertext vector. The key matrix is a square matrix that is used for both encryption and decryption.

**Implementation Details**

This implementation of the Hill Cipher is written in Go and provides functions for encrypting and decrypting messages using the Hill Cipher algorithm. The implementation includes the following features:

- Key matrix generation: The implementation provides a function for generating a key matrix from a given string.
- Encryption: The implementation provides a function for encrypting a plaintext message using the Hill Cipher algorithm.
- Decryption: The implementation provides a function for decrypting a ciphertext message using the Hill Cipher algorithm.
- Testing: The implementation includes a Test suite to ensure that the functions are working correctly.

**Usage**

To use the Hill Cipher implementation, you can call the `Encrypt` and `Decrypt` functions, passing in the plaintext or ciphertext message and the key matrix.

```
func main() {
    // Generate a key matrix from a given string
    keyMatrix, err := GetKeyMatrix("ABCDEFGHI")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Encrypt a plaintext message using the Hill Cipher algorithm
    plaintext := "ACT"
    encrypted, err := Encrypt(plaintext, "GYBNQKURP")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Decrypt a ciphertext message using the Hill Cipher algorithm
    decrypted, err := Decrypt(encrypted, "GYBNQKURP")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Plaintext:", plaintext)
    fmt.Println("Encrypted:", encrypted)
    fmt.Println("Decrypted:", decrypted)
}
```

**Testing**

The implementation includes a comprehensive test suite to ensure that the functions are working correctly. You can run the tests using the following command:

```
go test
```

**Lore**

The Hill Cipher has a rich history, dating back to the 1920s. It was invented by Lester S. Hill, an American mathematician and cryptographer. The cipher was designed to be a more secure alternative to traditional substitution ciphers, and it was widely used during World War II.

Despite its security advantages, the Hill Cipher has some limitations. It is vulnerable to frequency analysis attacks, and it can be broken using a known-plaintext attack.

In recent years, the Hill Cipher has seen a resurgence in popularity, particularly among cryptography enthusiasts and historians. It remains an important part of cryptography's history and continues to be studied and used by cryptographers around the world.
