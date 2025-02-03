# SHA-256 Implementation in Go

This repository contains a pure Go implementation of the SHA-256 cryptographic hash algorithm. SHA-256 is part of the SHA-2 family of cryptographic hash functions designed by the National Security Agency (NSA). It is widely used in various security applications and protocols, including TLS, SSL, PGP, SSH, and Bitcoin.

This implementation is designed to be educational, transparent, and easy to understand. It includes a step-by-step breakdown of the SHA-256 algorithm, including message padding, message schedule expansion, and the compression function.

---

## Table of Contents

1. [Lore](#lore)
2. [Features](#features)
3. [How It Works](#how-it-works)
4. [Usage](#usage)
5. [Testing](#testing)

---

## Lore

The SHA-256 algorithm is a cornerstone of modern cryptography. It was first published in 2001 as part of the SHA-2 family, designed to replace the older SHA-1 algorithm, which had begun to show vulnerabilities. SHA-256 produces a 256-bit (32-byte) hash value, typically rendered as a 64-character hexadecimal number.

This implementation was created as a learning exercise to understand the inner workings of SHA-256. By breaking down the algorithm into its fundamental components—message padding, message schedule expansion, and the compression function—this project aims to demystify the process of cryptographic hashing.

---

## Features

- **Pure Go Implementation**: No external dependencies; written entirely in Go.
- **Educational Focus**: Detailed comments and intermediate state logging to help understand the algorithm.

---

## How It Works

### SHA-256 Algorithm Overview

1. **Message Padding**:
   - The input message is padded to ensure its length is congruent to 448 modulo 512.
   - A `1` bit is appended, followed by `0` bits, and finally the original message length in bits.

2. **Message Schedule Expansion**:
   - The padded message is divided into 512-bit blocks.
   - Each block is expanded into 64 32-bit words using a specific algorithm.

3. **Compression Function**:
   - The expanded message block is processed using a series of logical operations (e.g., bitwise shifts, rotations, and additions).
   - The compression function updates the hash state (eight 32-bit words) for each block.

4. **Final Hash**:
   - After processing all blocks, the final hash value is computed by concatenating the eight 32-bit words.

### Key Components

- **Initial Hash Values**: Generated from the fractional parts of the square roots of the first 8 prime numbers.
- **Round Constants**: Generated from the fractional parts of the cube roots of the first 64 prime numbers.
- **Bitwise Operations**: Includes right rotations, shifts, and logical functions (e.g., Ch, Maj).

---

## Usage

### Basic Usage

The `sha256` function can be used to compute the hash of a message:

```go

func main() {
	message := "hello world"
	hash, _ := sha256.sha256([]byte(message))
	fmt.Printf("Input: %s\nHash: %x\n", message, hash)
}
```

### Intermediate States

The `sha256` function also returns intermediate states for debugging and educational purposes:

```go
hash, intermediateStates := sha256.sha256([]byte(message))
for i, state := range intermediateStates {
	fmt.Printf("Block %d: %08x %08x %08x %08x %08x %08x %08x %08x\n",
		i+1, state[0], state[1], state[2], state[3], state[4], state[5], state[6], state[7])
}
```

---

## Testing

The implementation includes test cases. To run the tests:

```bash
go test -v
```

### Example Test Output

```
=== RUN   TestSHA256
Input: ""
Expected: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
Actual:   e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
--- PASS: TestSHA256 (0.00s)
```

---

## Acknowledgments

- The [NIST FIPS 180-4](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.180-4.pdf) specification for SHA-256.
- The Go programming language community for their excellent documentation and tools.

---

Happy hashing! 🚀
