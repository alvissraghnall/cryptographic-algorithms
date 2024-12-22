package main

import (
	"errors"
//	"fmt"
	"math"
)

type Matrix [][]rune
type Vector []rune

func gcd (a, b rune) rune {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (matrix Matrix) IsMatrixInvertible () (bool, error) {
	if len(matrix) != len(matrix[0]) {
		return false, errors.New("Matrix not square!")
	}

	determinant := matrix.CalculateDeterminant() % 26
	if determinant < 0 {
		determinant += 26
	}
	/*
	if len(matrix) == 2 {
		return (determinant != 0 && gcd(determinant, 26) == 1), nil
	}
	*/

	if determinant % 2 == 0 || determinant % 13 == 0 {
		return false, nil
	}
	return true, nil
}

func (matrix Matrix) CalculateDeterminant() rune {
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	if len(matrix) == 2 {
		return (matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0])
	}

	var determinant rune = 0
	for i := 0; i < len(matrix); i++ {
		/// Get Minor here
		minor := getMinor(matrix, 0, rune(i))
		determinant += rune(math.Pow(-1, float64(i))) * matrix[0][i] * minor.CalculateDeterminant()
	}
	return determinant
}

func getMinor(matrix Matrix, row, col rune) Matrix {
	minor := make(Matrix, len(matrix)-1)
	for i := range minor {
		minor[i] = make([]rune, len(matrix[0])-1)
	}

	for i := 0; i < len(matrix); i++ {
		// if first row, skip, according to formula
		if i == int(row) {
			continue
		}

		//iterate over columns
		for j := 0; j < len(matrix[0]); j++ {
			// skip the current column being worked on.
			if j == int(col) {
				continue
			}

			newRow := i
			if i > int(row) {
				newRow--
			}

			newCol := j
			if j > int(col) {
				newCol--
			}

			minor[newRow][newCol] = matrix[i][j]
		}

	}

	return minor
}

func IsPerfectSquare (x float64) bool {
	if x < 0 {
		return false
	}
	sr := math.Sqrt(x)
	return math.Floor(sr)*math.Floor(sr) == x
}

// MatrixVectorMultiplication multiplies a matrix and a vector
func MatrixVectorMultiplication(matrix Matrix, vector Vector) (Vector, error) {
	// Check if matrix and vector can be multiplied
	if len(matrix[0]) != len(vector) {
		return nil, errors.New("matrix and vector cannot be multiplied")
	}

	// Create result vector
	result := make(Vector, len(matrix))

	// Multiply matrix and vector
	for i := range matrix {
		var sum rune = 0
		for j := range matrix[0] {
			sum += matrix[i][j] * vector[j]
		}
//		fmt.Println(sum)
		result[i] = sum % 26
	}

	return result, nil
}


func (matrix Matrix) Adjugate() Matrix {
        adjugate := make(Matrix, len(matrix))
        for i := range adjugate {
                adjugate[i] = make([]rune, len(matrix[0]))
        }

        for i := 0; i < len(matrix); i++ {
                for j := 0; j < len(matrix[0]); j++ {
                        minor := getMinor(matrix, rune(i), rune(j))
                        det := minor.CalculateDeterminant()
                        adjugate[i][j] = rune(int(det) * int(math.Pow(-1, float64(i+j))))
                }
        }

        return transpose(adjugate)
}

func transpose(matrix Matrix) Matrix {
        transposed := make(Matrix, len(matrix[0]))
        for i := range transposed {
                transposed[i] = make([]rune, len(matrix))
        }

        for i := 0; i < len(matrix); i++ {
                for j := 0; j < len(matrix[0]); j++ {
                        transposed[j][i] = matrix[i][j]
                }
        }

        return transposed
}
