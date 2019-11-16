package main

import (
	"fmt"
	"strconv"
)

func jumpingOnClouds(c []int32) int32 {

	var result int32
	var i int

	for i < len(c)-1 {
		if i+2 < len(c) && c[i+2] == 0 {
			i = i + 2
		} else {
			i++
		}
		result++
		println("POS: " + strconv.Itoa(i) + "result: " + fmt.Sprint(result))
	}

	return result
}

func main_CLOUDS() {

	result := jumpingOnClouds([]int32{0})

	println(result)
}
