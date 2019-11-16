package main

import (
	"fmt"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {

	magazineMap := make(map[string]int)
	noteMap := make(map[string]int)

	for i := range magazine {
		magazineMap[magazine[i]]++
	}

	for i := range note {
		noteMap[note[i]]++
	}
	for word, count := range noteMap {
		if count > magazineMap[word] {
			fmt.Println("No")
			break
		}
	}

	// for i := range note {

	// 	if !(magazineMap[note[i]]) {
	// 		fmt.Print("No")
	// 		return
	// 	}
	// 	magazineMap[note[i]] = false

	// }
	fmt.Println("Yes")
}

func printMap(myMap map[string]bool) {
	for key, value := range myMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}

func main() {

	checkMagazine([]string{""}, []string{""})
}

// THIS FAILS:

// 15 17
// o l x imjaw bee khmla v o v o imjaw l khmla imjaw x
// imjaw l khmla x imjaw o l l o khmla v bee o o imjaw imjaw

// EXPECTED NO
