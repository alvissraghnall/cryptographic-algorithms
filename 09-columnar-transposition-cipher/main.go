package main

import (
  "fmt"
)

func main () {
  e, err := encrypt("WE ARE DISCOVERED FLEE AT ONCE", "Zebra")
  if err != nil {
      fmt.Println(err)
  }
  fmt.Println(e)
}
