
# MAD Hash Function

The MAD hash function is a simple and efficient hash function designed for general-purpose use. It utilizes the FNV-1a hash function to sum the input string into a 64 bit unsigned integer, and uses a combination of multiplication, addition, and division to produce a hash value.

### Features

- Simple and efficient implementation
- Utilizes on the FNV-1a hash function
- Uses a combination of multiplication, addition, and division to produce a hash value

### Usage

To use the MAD hash function, simply call the `madHash` function and pass in the string you want to hash. The function will return a hash value as an integer.

```
hashValue := madHash("test")
```

### Testing

The MAD hash function has been tested for correctness and performance using a variety of test cases. The tests can be run using the `go test` command.

### Dependencies

The MAD hash function has no external dependencies and can be used as a standalone library.
