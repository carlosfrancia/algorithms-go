package main

import (
	"fmt"
)

// Complete the gameOfThrones function below.
func gameOfThrones(s string) string {

	ocurrences := make(map[string]int)
	foundSingle := false

	for _, value := range s {
		ocurrences[string(value)]++
	}
	printMapStringInt(ocurrences)
	for _, value := range ocurrences {

		if value%2 != 0 && foundSingle {
			return ("NO")
			break
		}
		if value%2 != 0 && !foundSingle {
			foundSingle = true
		}
	}
	return "YES"
}

func mainGOT() {

	fmt.Println(gameOfThrones("aaabbbb"))

}
