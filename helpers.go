package main

import (
	"fmt"
	"strconv"
)

func printMapStringBool(myMap map[string]bool) {
	for key, value := range myMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}

func printMapStringInt(myMap map[string]int) {
	for key, value := range myMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}

func printArrayInt32(arr []int32) {
	for i := range arr {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func printMapIntBool(myMap map[int]bool) {
	for key, value := range myMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}
func printArrayString(arr []string) {
	for i := range arr {
		fmt.Print(arr[i])
	}
	fmt.Println()
}

func printMatrixInt32(arr [][]int32) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

func isInt(n string) bool {

	_, err := strconv.Atoi(n)
	if err == nil {
		return true
	}
	return false
}
