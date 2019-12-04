package main

import (
	"fmt"
)

// Complete the countSwaps function below.
func countSwaps(a []int32) []int32 {
	var number int

	for i := len(a); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if a[j] > a[j+1] {
				toSwap := a[j]
				a[j] = a[j+1]
				a[j+1] = toSwap
				number++
			}
		}
	}

	// for i := 0; i < len(a); i++ {
	// 	for j := 0; j < (len(a) - 1); j++ {
	// 		if a[j] > a[j+1] {
	// 			toSwap := a[j]
	// 			a[j] = a[j+1]
	// 			a[j+1] = toSwap
	// 			number++
	// 		}
	// 	}
	// }
	fmt.Println("Array is sorted in", number, "swaps.")
	fmt.Println("First Element:", a[0])
	fmt.Println("Last Element:", a[len(a)-1])
	return a
}

func main_ARRAY_SORT() {
	countSwaps([]int32{30, 1, 20, 100})

	printArrayInt32(countSwaps([]int32{-11, -30, -1, -20, -100, 0}))
}
