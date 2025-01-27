package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
//  md5([]byte("+#+#+"))
//  md5([]byte("SAM SJSNENWJWJHSBSHSHSHSHSHWBSJAOKAjjwjsjjs722eijejdjsjsjsjsj"))
//  md5([]byte {0x80, 0x57, 0x9A, 0x8b, 0x6d, 0x27,0x08})

  hash := md5([]byte(""))
  fmt.Printf("%x\n", md5_buf([]byte("")))
  fmt.Printf("%x\n", hash)
  fmt.Println(fmt.Sprintf("%x", hash) == "d41d8cd98f00b204e9800998ecf8427e")
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

  // Initialize variables:
  var a0, b0, c0, d0 uint32 = 0x67452301, 0xefcdab89 , 0x98badcfe, 0x10325476   // A
  
  var i uint32
  var K [64]uint32
  shift := [...]uint{7, 12, 17, 22, 5, 9, 14, 20, 4, 11, 16, 23, 6, 10, 15, 21}

  for i = range 64 {
    K[i] = uint32(math.Floor((1 << 32) * math.Abs(math.Sin(float64(i) + 1))))
  }

  // process message in 512-bit chunks:
  for k := 0; k < len(paddedMessage); k += 64 {
    chunk := paddedMessage[k:min(k+64, len(paddedMessage))]
    // fmt.Println(string(chunk))
    M := make([]uint32, 16)

    for j := range len(M) {
      // fmt.Println(j, M, j * 4, chunk[(j * 4) + 3])
      M[j] = uint32(chunk[j * 4]) | uint32(chunk[j*4+1] << 8) | uint32(chunk[j*4+2] << 16) | uint32(chunk[j*4+3] << 24)
    }

    A, B, C, D := a0, b0, c0, d0

    for i = range 64 {

      round := i >> 4
      var F, g uint32 
      switch {
      case i >= 0 && i <= 15:
        F = ( B & C ) | (^B & D)
        g = i
      
      case i >= 16 && i <= 31:
        F = (D & B ) | (^D & C)
        g = ((5 * i) + 1) & 0x0F

      case i >= 32 && i <= 47:
        F = B ^ C ^ D 
        g = (3 * i + 5) & 0x0F

      case i >= 48 && i <= 63:
        F = C ^ (B | ^D )
        g = (7 * i)% 16
      }

      F += A + K[i] + M[g]

      sa := shift[(round << 2) | (i & 3)]

      A, D, C, B = D, C, B, leftShift(F, sa) + B 

    }

    a0 += A
    b0 += B
    c0 += C
    d0 += D
  }
  
  fmt.Println(len(paddedMessage) * 8)

  var digest bytes.Buffer

  binary.Write(&digest, binary.LittleEndian, []uint32{a0, b0, c0, d0})
  digestByteArr := digest.Bytes()

//  return [16]byte(digestByteArr)

  var result [16]byte
  as := [...]uint32{a0, b0, c0, d0}
  for i, a := range as {
          result[i*4] = byte(a)
          result[i*4+1] = byte(a >> 8)
          result[i*4+2] = byte(a >> 16)
          result[i*4+3] = byte(a >> 24)
  }

  fmt.Println(digestByteArr, result)
  return result

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


func leftShift(value uint32, shift uint) uint32 {
  return ((value << shift) | (value >> (32 - shift))) // & 0xFFFFFFFF
}
