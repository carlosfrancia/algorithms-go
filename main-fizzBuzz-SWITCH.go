package main

import (
	"fmt"
)

func fizzBuzzasdfasdf(n int32) {
	var i int32
	//print(i % 3)
	for i = 1; i <= n; i++ {
		mod3 := i % 3
		mod5 := i % 5
		if (mod3 == 0) && (mod5 == 0) {
			fmt.Println("FizzBuzz")
		} else {
			if mod3 == 0 {
				fmt.Println("Fizz")
			} else {
				if mod5 == 0 {
					fmt.Println("Buzz")
				} else {
					fmt.Println(i)
				}
			}
		}
	}

}

func mainasdasdf() {

}
