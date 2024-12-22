package main

import (
  "fmt"
)

func main () {
  fmt.Println(createKeyGrid("MONARCHY"))
  plaintext := "hello"
    digraphs := splitIntoDigraphs(plaintext)
    for _, digraph := range digraphs {
    fmt.Println(digraph)
  }
}
