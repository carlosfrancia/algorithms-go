package main

import (
	"fmt"
)

func maxDifference(arr []int32) int32 {
	// Write your code here
	if len(arr) == 1 {
		return (arr[0])
	}
	result := arr[1] - arr[0]
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] >= min {
			if arr[i]-min > result {
				result = arr[i] - min
			}
		}

		if arr[i] < min {
			min = arr[i]
		}

	}

	return result
}

func main() {

	maxDifference([]int32{2, 3, 4})
	c := 5 << 2
	x := 5 >> 2
	fmt.Println(c, x)
}
