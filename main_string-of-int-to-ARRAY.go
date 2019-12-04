package main

import (
	"fmt"
	"strconv"
	"strings"
)

func smallest2(n string) {
	m := make(map[int]bool)
	s := strings.Split(n, ",")
	isInvalid := false
	printArrayString(s)
	for i := range s {
		fmt.Println("I: ", i)
		//fmt.Println("V: ", value)
		fmt.Println("[i]: ", s[i])
		if !isInt2(string(s[i])) {
			fmt.Print("Invalid")
			isInvalid = true
			break
		} else {
			num, _ := strconv.Atoi(string(s[i]))
			if num > 1000 {
				fmt.Print("Invalid")
				isInvalid = true
				break
			} else {
				m[num] = true
			}
		}
	}
	if !isInvalid {
		for i := 1; i < 1001; i++ {
			if !m[i] {
				fmt.Print(i)
				break
			}
		}
	}

}

func isInt2(n string) bool {
	_, err := strconv.Atoi(n)
	if err == nil {
		return true
	}
	return false
}

func main2() {

	smallest("1,3,6,4,6,7,8")

}
