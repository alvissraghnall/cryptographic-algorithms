package main 

import "strings"


func main () { 
}

func hashAscii (text string) int {
  HASH_TABLE_SIZE := 1024

  text = strings.ReplaceAll(text, " ", "")

  hashValue := 0
  for _, v := range text {
    hashValue += int(v)

  }
  return hashValue % HASH_TABLE_SIZE
}

func hashBinaryShift (text string) int {
  var HASH_TABLE_SIZE int64 = 1024

  text = strings.ReplaceAll(text, " ", "")

  var hashValue int64 = 0
  for _, v := range text {
    hashValue = ((hashValue << 5) + int64(v)) % HASH_TABLE_SIZE
  }

  return int(hashValue % HASH_TABLE_SIZE)
}
