package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
  md5([]byte("+#+#+"))
  md5([]byte("SAM SJSNENWJWJHSBSHSHSHSHSHWBSJAOKAjjwjsjjs722eijejdjsjsjsjsj"))
  md5([]byte {0x80, 0x57, 0x9A, 0x8b, 0x6d, 0x27,0x08})
}

func md5 (message []byte) [16]byte {
  // APPEND PADDING BITS
  paddedMessage := append(message, 0x80)

  padLength := 56 - (len(paddedMessage) % 64)

  if padLength < 0 {
    padLength += 64
  }
  for i := 0; i < padLength; i++ {
    paddedMessage = append(paddedMessage, 0x00)
  }

  // APPEND LENGTH H
  messageLengthInBits := uint64(len(message)) * 8
  // paddedMessage = append(paddedMessage, swapUint64(uint64(messageLengthInBits & 0xFFFFFFFFFFFFFFFF)))
  //
  var bytesToWrite [8]byte
  binary.LittleEndian.PutUint64(bytesToWrite[:], messageLengthInBits)
  paddedMessage = append(paddedMessage, bytesToWrite[:]...) 

  // process message in 512-bit chunks:
  for i := 0; i < len(paddedMessage); i += 64 {
     chunk := paddedMessage[i:min(i+64, len(paddedMessage))]
    M := make([]uint32, 16)

    for j := range M {
      M[j] = uint32(chunk[j * 4]) | uint32(chunk[j*4+1] << 8) | uint32(chunk[j*4+2] << 16) | uint32(chunk[j*4+3] << 24)
    }
  }
  
  fmt.Println(len(paddedMessage) * 8)

  return [16]byte{}
}

func swapUint64(val uint64) uint64 {
	val = ((val << 8) & 0xFF00FF00FF00FF00) | ((val >> 8) & 0x00FF00FF00FF00FF)
	val = ((val << 16) & 0xFFFF0000FFFF0000) | ((val >> 16) & 0x0000FFFF0000FFFF)
	return (val << 32) | (val >> 32)
} 


func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}
