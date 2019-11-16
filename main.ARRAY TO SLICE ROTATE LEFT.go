package main

import (
	"fmt"
)

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {

	var i int32
	slice := a[:]
	for i = 0; i < d; i++ {
		slice = append(slice, slice[0])
		fmt.Print("Append ")
		printArray(slice)
		slice = slice[1:] // remove first
		fmt.Print("Remove ")
		printArray(slice)
	}
	return slice
}

func main_ARRAY_ROTATE() {

	var arr []int32

	arr = []int32{1, 2, 3, 4, 5}

	//printMatrix(matrix)

	result := rotLeft(arr, 4)
	printArray(result)
}

func printArray(arr []int32) {
	for i := range arr {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
