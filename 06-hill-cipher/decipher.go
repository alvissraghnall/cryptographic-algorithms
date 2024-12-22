package main

import (
	"errors"
	"strconv"
)

func (matrix Matrix) ScalarMultiply(scalar int, mod int) Matrix {
	result := make(Matrix, len(matrix))
	for i := range result {
		result[i] = make([]rune, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			result[i][j] = rune((int(matrix[i][j]) * scalar) % mod)
		}
	}

	return result
}

func (matrix Matrix) ModInverse(mod int) (Matrix, error) {
	det := matrix.CalculateDeterminant()
	if det == 0 {
		return nil, errors.New("matrix is not invertible")
	}

	detInv := modInverse(int(det), mod)
	if detInv == 0 {
		return nil, errors.New("matrix is not invertible modulo " + strconv.Itoa(mod))
	}

	adjugate := matrix.Adjugate()
	adjugate = adjugate.ScalarMultiply(detInv, mod)

	// Add 26 to any negative values
	for i := range adjugate {
		for j := range adjugate[0] {
			if adjugate[i][j] < 0 {
				adjugate[i][j] += 26
			}
		}
	}

	return adjugate, nil
}

func modInverse(a, m int) int {
	m0 := m
	y := 0
	x := 1

	if m == 1 {
		return 0
	}

	for a > 1 {
		q := a / m
		t := m

		m = a % m
		a = t
		t = y

		y = x - q*y
		x = t
	}

	if x < 0 {
		x += m0
	}

	return x
}


func Decrypt (text string, key string) (string, error) {

  keyMatrix, err := GetKeyMatrix(key)
  if err != nil {
    return "", err
  }

  cipherTextVector, err := ParseTextToVector(text, len(keyMatrix))
  if err != nil {
    return "", err
  }

  inverseKeyMatrix, err := keyMatrix.ModInverse(26)
  if err != nil {
    return "", err
  }

  plainTextVector := make([][]rune, len(cipherTextVector))

  for c := range plainTextVector {
    plainTextVector[c] = make([]rune, len(cipherTextVector[0]))
  }

  for i, v := range cipherTextVector {
    multiplied, err := MatrixVectorMultiplication(inverseKeyMatrix, v)

    if err != nil {
      return "", err
    }
    plainTextVector[i] = multiplied
  }

  result := ""

  for _, col := range plainTextVector {
    for _, v := range col {
      result += string(v + 65)
    }
  }

  return result, nil
}
