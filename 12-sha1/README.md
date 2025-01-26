# SHA1 Implementation in Go

### Background

The Secure Hash Algorithm 1 (SHA1) is a cryptographic hash function that was designed by the United States National Security Agency (NSA) in 1995. Although it was once widely used, SHA1 has been largely deprecated due to security concerns and the discovery of collisions.

### This Implementation

This implementation of SHA1 in Go is a faithful reproduction of the original algorithm. It was created as a learning exercise and to provide a reference implementation for educational purposes.

### Features

- Implements the full SHA1 algorithm, including padding and hashing
- Supports inputs of any length
- Returns a 20-byte hash value
- Includes a comprehensive test suite to ensure correctness

### Performance

This implementation is not optimized for performance and is intended for educational purposes only. For production use, a more optimized implementation should be used.

### Security

As mentioned earlier, SHA1 is no longer considered secure for cryptographic purposes. This implementation should not be used for any security-critical applications.

### Testing

The test suite includes a range of tests to ensure correctness, including:

- Test vectors from the SHA1 specification
- Edge cases, such as empty inputs and inputs of length 1
- Consistency tests to ensure that multiple calls with the same input produce the same output

### Benchmarking

The implementation includes two benchmarking tests to measure performance:

- `BenchmarkSHA1`: measures the performance of hashing a small input
- `BenchmarkSHA1Large`: measures the performance of hashing a large input (1MB)
