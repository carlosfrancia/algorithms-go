package main

import (
	"fmt"
	"strconv"
	"strings"
)

func shopping(n string) {
	result := 0
	budget := 10000
	s := strings.Split(n, ",")
	printArrayString(s)
	for i := range s {
		if s[i] != "" {
			if !isInt(s[i]) {
				fmt.Println("Invalid", string(s[i]))
				break
			} else {
				price, _ := strconv.Atoi(s[i])
				budget = budget - price
				if budget < 0 {
					fmt.Print(result)
					return
				}
				result++
			}
		}

	}
	fmt.Print(result)
}

func main_SHOPPING() {

	shopping("9850,100,")
	//FIX THIS ^^ Should return 2

}
