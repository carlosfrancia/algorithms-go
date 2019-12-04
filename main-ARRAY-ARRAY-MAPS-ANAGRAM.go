package main

import (
	"fmt"
)

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	var r int32
	mapS1 := make(map[rune]int32)
	for _, letter := range a {

		mapS1[letter]++
	}
	mapS2 := make(map[rune]int32)
	for _, letter := range b {

		mapS2[letter]++
	}
	for key := range mapS1 {
		if mapS1[key] > mapS2[key] {
			r = r + (mapS1[key] - mapS2[key])
		}
	}
	for key := range mapS2 {
		if mapS2[key] > mapS1[key] {
			r = r + (mapS2[key] - mapS1[key])
		}
	}
	return r
}

func main_MAPSANA() {
	fmt.Print(makeAnagram("cde", "abc"))
}
