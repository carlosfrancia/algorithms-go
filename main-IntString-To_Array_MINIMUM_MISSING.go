package main

import (
	"fmt"
	"strconv"
	"strings"
)

func smallest(n string) {
	m := make(map[int]bool)
	s := strings.Split(n, ",")
	//isInvalid := false
	printArrayString(s)
	for i := range s {
		if s[i] != "" {
			if !isInt(s[i]) {
				fmt.Println("NUMBER", s[i])
				fmt.Println("Invalid")
				// isInvalid = true
				return
			} else {
				num, _ := strconv.Atoi(s[i])
				if num > 1000 {
					fmt.Print("Invalid", num)
					// isInvalid = true
					return
				}
				m[num] = true
			}
		}

	}
	for i := 1; i < 1001; i++ {
		if !m[i] {
			fmt.Print(i)
			break
		}
	}

	// if !isInvalid {
	// 	for i := 1; i < 1001; i++ {
	// 		if !m[i] {
	// 			fmt.Print(i)
	// 			break
	// 		}
	// 	}
	// }
}

func main_MINIMUM() {

	smallest("1,3,6,4,6,7,8,")

}
