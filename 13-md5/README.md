# Pure Go MD5 Implementation

A straightforward implementation of the MD5 (Message Digest Algorithm 5) hashing algorithm in Go. This implementation prioritizes readability and understanding over performance, making it ideal for learning about hash functions and cryptographic primitives.

## About MD5

MD5 was designed by Ronald Rivest in 1991 to succeed MD4, and was published in 1992 as RFC 1321. The algorithm takes an input message of arbitrary length and produces a 128-bit (16-byte) hash value, traditionally expressed as a 32-digit hexadecimal number.

### Historical Significance

MD5 played a crucial role in the development of cryptographic hash functions and was widely used throughout the 1990s and early 2000s. Some interesting historical facts:

- The "MD" in MD5 stands for "Message Digest"
- It was extensively used for file integrity verification and password hashing
- The algorithm was developed at MIT by Ronald Rivest, the "R" in RSA
- Each round of MD5 was carefully designed to create an "avalanche effect" where small changes in input create large changes in output

### Security Status

**Important**: This implementation is for educational purposes only. MD5 is cryptographically broken and should not be used for any secure hashing.

- In 1996, a flaw was found that while not fatal, suggested MD5 was not collision-resistant
- In 2004, Wang et al. demonstrated the first collision attack
- By 2012, the Flame malware exploited MD5 weaknesses to fake Microsoft digital signatures

## Implementation Details

This implementation follows RFC 1321 and includes:

- Message padding according to the MD5 specification
- Processing of message in 512-bit blocks
- Four-round processing with 16 operations per round
- Little-endian byte ordering throughout

### Usage

```go
package main

import (
    "fmt"
)

func main() {
    message := []byte("Hello, World!")
    hash := md5(message)
    fmt.Printf("%x\n", hash)
}
```

### Key Features

- Pure Go implementation with no external dependencies
- Clear, commented code explaining each step
- Follows the original MD5 specification precisely
- Handles messages of any length
- Uses standard library's binary package for endian operations

## Technical Details

The MD5 algorithm operates in these steps:

1. **Padding**: Append bits to the message so its length is congruent to 448 modulo 512
2. **Length**: Append a 64-bit representation of the original message length
3. **Initialization**: Set up four 32-bit registers (A, B, C, D)
4. **Processing**: Process message in 512-bit chunks through four rounds of operations
5. **Output**: Concatenate the final values of A, B, C, D to produce the hash

### The Four Rounds

Each round uses a different nonlinear function:
```
Round 1: F(B,C,D) = (B ∧ C) ∨ (¬B ∧ D)
Round 2: G(B,C,D) = (B ∧ D) ∨ (C ∧ ¬D)
Round 3: H(B,C,D) = B ⊕ C ⊕ D
Round 4: I(B,C,D) = C ⊕ (B ∨ ¬D)
```

## Fun Facts

- MD5 was used to index faces in the early days of Google Images
- The "avalanche effect" is so strong that changing a single bit in the input typically changes about half the bits in the output
