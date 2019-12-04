package main

import (
	"fmt"
)

// Complete the alternatingCharacters function below.
func alternatingCharacters2(s string) int32 {
	var r int32

	//	for i, let := range s {
	for i := 0; i < len(s)-1; i++ {
		if i < len(s)-1 {
			if string(s[i]) == string(s[i+1]) {
				fmt.Print(string(s[i+1]))
				r++
			}
		}
	}

	return r

}

func main_STRING_ITERATE() {

	result := alternatingCharacters2("AAABBB")
	println(result)
}
