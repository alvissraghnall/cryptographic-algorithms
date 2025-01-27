package main

import (
        // "bytes"
        "bytes"
        "encoding/binary"
        "fmt"
        "math"
)

func md5_buf (message []byte) [16]byte {

  paddedMessage := bytes.NewBuffer(message)
  // APPEND PADDING BITS
  paddedMessage.WriteByte(0x80)

  padLength := 56 - paddedMessage.Len() % 64

  if padLength < 0 {
    padLength += 64
  }
  for i := 0; i < padLength; i++ {
    paddedMessage.WriteByte(0x00)
  }

  // APPEND LENGTH H
  messageLengthInBits := uint64(len(message)) * 8
  // paddedMessage = append(paddedMessage, swapUint64(uint64(messageLengthInBits & 0xFFFFFFFFFFFFFFFF)))
  //
  binary.Write(paddedMessage, binary.LittleEndian, messageLengthInBits)

  // fmt.Println(paddedMessage.Bytes())
  // Initialize variables:
  var a0, b0, c0, d0 uint32 = 0x67452301, 0xefcdab89 , 0x98badcfe, 0x10325476   // A

  var i uint32
  var K [64]uint32
  shift := [...]uint{7, 12, 17, 22, 5, 9, 14, 20, 4, 11, 16, 23, 6, 10, 15, 21}

  for i = range 64 {
    K[i] = uint32(math.Floor((1 << 32) * math.Abs(math.Sin(float64(i) + 1))))
  }

  // process message in 512-bit chunks:
  var bytesToRead [16]uint32
  for binary.Read(paddedMessage, binary.LittleEndian, bytesToRead[:]) == nil {
  //  chunk := paddedMessage[k:min(k+64, len(paddedMessage))]
  // M := make([]uint32, 16)

    A, B, C, D := a0, b0, c0, d0

    for i = range 64 {

      var F, g uint32
      round := i >> 4
      switch {
      case i >= 0 && i <= 15:
        F = ( B & C ) | (^B & D)
        g = i

      case i >= 15 && i <= 31:
        F = (D & B ) | (^D & C)
        g = ((5 * i) + 1) & 0x0F

      case i >= 32 && i <= 47:
        F = B ^ C ^ D
        g = (3 * i + 5) & 0x0F

      case i >= 48 && i <= 63:
        F = C ^ (B | ^D )
        g = (7 * i) % 16
      }

      F += A + K[i] + bytesToRead[g]

      sa := shift[(round << 2) | (i & 3)]
      // fmt.Println(sa)

      A, D, C, B = D, C, B, leftShift(F, sa) + B

    }

    a0 += A
    b0 += B
    c0 += C
    d0 += D

    fmt.Println(a0,b0,c0,d0)
  }

  var digest bytes.Buffer

  binary.Write(&digest, binary.LittleEndian, []uint32{a0, b0, c0, d0})
  digestByteArr := digest.Bytes()

  return [16]byte(digestByteArr)

}

