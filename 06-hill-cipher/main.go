package main

import (
	"fmt"
)

func main () {
	matrix, err := GetKeyMatrix("ABCDEFGHI")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matrix)
	
        matrixTwo, errTwo := GetKeyMatrix("ABCDEFHI")
        if errTwo != nil {
                fmt.Println(errTwo)
        }
        fmt.Println(matrixTwo)

	matrixThree, errThree := GetKeyMatrix("GYBNQKURP")
        if errThree != nil {
                fmt.Println(errThree)
        }
        fmt.Println(matrixThree)

	inverseMatrix, err := matrixThree.ModInverse(26)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inverseMatrix)

	vector, errFour := ParseTextToVector("ACT", 3)
	if errFour != nil {
		fmt.Println(errFour)
	}
	fmt.Println(vector)

	encryptedOne, err := Encrypt("HELP", "DDCF")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encryptedOne)

	encryptedTwo, err := Encrypt("ACT", "GYBNQKURP")
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(encryptedTwo)

	encryptedThree, err := Encrypt("CAT", "GYBNQKURP")
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(encryptedThree)

	decryptedOne, err := Decrypt("HIAT", "DDCF")
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(decryptedOne)

	decryptedTwo, err := Decrypt("POH", "GYBNQKURP")
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(decryptedTwo)

        decryptedThree, err := Decrypt("FIN", "GYBNQKURP")
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(decryptedThree)

}
