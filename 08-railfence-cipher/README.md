# Rail Fence Cipher

The Rail Fence Cipher is a classic encryption technique that has been used for centuries. It is also known as the Zigzag Cipher.

### Lore 

The **Rail Fence Cipher** is believed to have originated in ancient Greece, where it was used to send secret messages. The cipher was later used by the Romans and other civilizations to protect their communications.

The cipher gets its name from the way the plaintext is written in a zigzag pattern, resembling a rail fence. The ciphertext is then read off in rows, rather than columns, to create the encrypted message.

### How it Works

The Rail Fence Cipher is a transposition cipher, which means that it rearranges the letters of the plaintext to create the ciphertext. The cipher works as follows:

1. The plaintext is written in a zigzag pattern, with each row representing a "rail" of the fence.
2. The number of rails is determined by the key, which is typically a small integer.
3. The ciphertext is read off in rows, rather than columns, to create the encrypted message.

### Implementation

This implementation of the Rail Fence Cipher is written in Go. It includes two main functions: `Encrypt` and `Decrypt`.

- `Encrypt` takes a plaintext string and a key (the number of rails) as input, and returns the encrypted ciphertext.
- `Decrypt` takes a ciphertext string and a key (the number of rails) as input, and returns the decrypted plaintext.

The implementation also includes several helper functions, including `FormRails`, which creates the zigzag pattern of the rail fence, and `toUpperCase` and `removeSpaces`, which are used to preprocess the plaintext.

### Example Use Cases

Here is an example of how to use the `Encrypt` function:

```
encrypted, err := Encrypt("IAMSTARBOY", 3)
if err != nil {
    fmt.Println(err)
}
fmt.Println(encrypted)
```

This code encrypts the plaintext "IAMSTARBOY" using a key of 3, and prints the resulting ciphertext.

Similarly, here is an example of how to use the `Decrypt` function:

```
decrypted, err := Decrypt("LXFOPVEFRNHR", 3)
if err != nil {
    fmt.Println(err)
}
fmt.Println(decrypted)
```

This code decrypts the ciphertext "LXFOPVEFRNHR" using a key of 3, and prints the resulting plaintext.

### Testing

The implementation includes a comprehensive test suite, which covers a range of test cases, including:

- Encrypting and decrypting plaintexts of varying lengths
- Using different keys (numbers of rails)
- Testing edge cases, such as empty plaintexts and invalid keys

### Contributing

Contributions are welcome. 
If you'd like to contribute to this implementation, please fork the repository and submit a pull request.

