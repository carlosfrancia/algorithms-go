package main

import (
	"fmt"
)

// Complete the countSwpricesps function below.
func maximumToys(prices []int32, k int32) int32 {
	var result int32

	for i := 0; i < len(prices); i++ {
		for j := len(prices) - 1; j > i; j-- {
			fmt.Println("j", prices[j])
			fmt.Println("j-1", prices[j-1])
			if prices[j] < prices[j-1] {
				toSwap := prices[j]
				prices[j] = prices[j-1]
				prices[j-1] = toSwap
			}
		}
		fmt.Println("First: ", prices[i])
		if prices[i] < k {
			result++
			k = k - prices[i]
		} else {
			return result
		}
	}

	return result
}

func main_ARRAY_SORT_INVERSE() {
	fmt.Print(maximumToys([]int32{1, 12, 5, 111, 200, 1000, 10}, 50))

}
