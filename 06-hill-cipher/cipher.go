package main

import (
  "math"
  "errors"
)

func GetKeyMatrix (key string) (Matrix, error) {
  if !IsPerfectSquare(float64(len(key))) {
    return nil, errors.New("Key must form an n x n matrix!")
  }
  n := int(math.Sqrt(float64(len(key))))
  keyMatrix := make(Matrix, n)

  for c := range keyMatrix {
    keyMatrix[c] = make([]rune, n)
  }

  k := 0
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      keyMatrix[i][j] = rune(key[k] % 65)
      k++
    }
  }

  return keyMatrix, nil
}

func ParseTextToVector (text string, keyMatrixLen int) ([][]rune, error) {

  if len(text) == 0 {
    return nil, errors.New("Text must contain at least one character!")
  }

  var textVector [][]rune

  if len(text) > keyMatrixLen {
    textVector = make([][]rune, int(math.Ceil(float64(len(text) / keyMatrixLen))))
  } else {
    textVector = make([][]rune, 1)
  }

  for c := range textVector {
    textVector[c] = make([]rune, keyMatrixLen)
  }

  for len(text) % keyMatrixLen != 0 {
    text += "X"
  }

  t := 0
  for i := 0; i < len(textVector); i++ {
    for j:= 0; j < keyMatrixLen; j++ {
      textVector[i][j] = rune(text[t] % 65)
      t++
    }
  }

  return textVector, nil
}

func Encrypt (text string, key string) (string, error) {

  keyMatrix, err := GetKeyMatrix(key)
  if err != nil {
    return "", err
  }

  isMatrixInvertible, err := keyMatrix.IsMatrixInvertible()

  if err != nil {
    return "", err
  }

  if !isMatrixInvertible {
    return "", errors.New("Key Matrix must be invertible to use the Hill Cipher!")
  }

  plainTextVector, err := ParseTextToVector(text, len(keyMatrix))
  if err != nil {
    return "", err
  }

  cipherTextVector := make([][]rune, len(plainTextVector))

  for c := range cipherTextVector {
    cipherTextVector[c] = make([]rune, len(plainTextVector[0]))
  }

  for i, v := range plainTextVector {
    multiplied, err := MatrixVectorMultiplication(keyMatrix, v) 

    if err != nil {
      return "", err
    }

    cipherTextVector[i] = multiplied
  }

  result := ""

  for _, col := range cipherTextVector {
    for _, v := range col {
      result += string(v + 65)
    }
  }

  return result, nil
}


