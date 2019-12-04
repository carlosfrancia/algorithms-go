package main

import (
	"fmt"
)

func mergeArrays(a []int32, b []int32) []int32 {
	// Write your code here

	var i, j, k int
	var result []int32
	slice := result[:]

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			slice = append(slice, a[i])
			i++
		} else {
			slice = append(slice, b[j])
			j++
		}
		k++
	}

	for i < len(a) {
		slice = append(slice, a[i])
		i++
		k++
	}

	for j < len(b) {
		slice = append(slice, b[j])
		j++
		k++
	}
	return slice
}

func main555() {
	// var n int
	// var a int
	// var b int
	// fmt.Scan(&n)

	// for i := 0; i < n; i++ {
	// 	var sum int
	// 	fmt.Scan(&a, &b)
	// 	sum = a + b
	// 	fmt.Println(uint(sum))
	// }

	fmt.Println("NN")
}

// for i := 2; i < len(arr); i++ {
// 	fmt.Println("INDEX: ", i)
// 	for j := 0; j < i; j++ {
// 		fmt.Println("i: ", arr[i])
// 		fmt.Println("j: ", arr[j])
// 		fmt.Println("REST: ", arr[i]-arr[j])
// 		dif := arr[i] - arr[j]
// 		if dif > result {
// 			result = dif
// 		}
// 	}
// }
