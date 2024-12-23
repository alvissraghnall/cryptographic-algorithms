package main

import (
  "fmt"
)

func main () {
  encrypted, err := Encrypt("GEEKSFORGEEKS", 3)
  if err != nil {
    fmt.Println(err)
  }
	fmt.Println(encrypted)
}
