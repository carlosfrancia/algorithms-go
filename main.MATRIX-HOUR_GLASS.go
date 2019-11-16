package main

import (
	"fmt"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var result int32

	for i := 0; i < 4; i++ {
		fmt.Print("i: ", i, "-")
		for j := 0; j < 4; j++ {
			if i == 0 && j == 0 {
				result = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			} else {
				acum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
				if result < acum {
					result = acum
				}
				fmt.Println("ACUM: ", acum)
			}
			fmt.Println("j: ", j)

		}
	}
	return result
}

func main_MATRIX() {

	var matrix [][]int32

	matrix = [][]int32{{-1, -1, 0, -9, -2, -2}, {-2, -1, -6, -8, -2, -5}, {-1, -1, -1, -2, -3, -4}, {-1, -9, -2, -4, -4, -5}, {-7, -3, -3, -2, -9, -9}, {-1, -3, -1, -2, -4, -5}}

	//printMatrix(matrix)

	result := hourglassSum(matrix)
	println(result)
}

func printMatrix(arr [][]int32) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
