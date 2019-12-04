package main

import (
	"fmt"
)

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {

	mapS1 := make(map[rune]bool)
	for _, letter := range s1 {

		mapS1[letter] = true
	}
	mapS2 := make(map[rune]bool)
	for _, letter := range s2 {

		mapS2[letter] = true
	}
	for key := range mapS2 {
		if mapS1[key] {
			return ("YES")
		}
	}
	return ("NO")
}

func main_ARRAY() {

	fmt.Print(twoStrings("hi", "world"))
}

// THIS FAILS:

// 15 17
// o l x imjaw bee khmla v o v o imjaw l khmla imjaw x
// imjaw l khmla x imjaw o l l o khmla v bee o o imjaw imjaw

// EXPECTED NO
