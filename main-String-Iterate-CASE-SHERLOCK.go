package main

import (
	"fmt"
)

// Complete the gameOfThrones function below.
func isValid(s string) string {

	var number, first, second, third int
	var firstK, secondK string
	i := 0
	ocurrences := make(map[string]int)
	foundSingle := false

	for _, value := range s {
		ocurrences[string(value)]++
	}
	printMapStringInt(ocurrences)

	for key, value := range ocurrences {
		if i == 0 {
			firstK = key
			first = value
		} else if i == 1 {
			secondK = key
			second = value
		} else if i == 2 {
			third = value
		} else {
			break
		}
		i++
	}

	diff := first - second
	if diff < -1 && first == 1 {
		number = second
		foundSingle = true
		delete(ocurrences, firstK)
	} else if diff > 1 && second == 1 {
		number = first
		foundSingle = true
		delete(ocurrences, secondK)
	} else if diff == 0 {
		number = first
	} else if diff == -1 {
		if first != 1 {
			number = first
		} else {
			if second == third {
				delete(ocurrences, firstK)
				fmt.Println("aferDelet")
				printMapStringInt(ocurrences)
				number = second
				foundSingle = true
			} else {
				return ("NO")
			}
		}
	} else if diff == 1 {
		if second != 1 {
			number = second
		} else {
			if first == third {
				delete(ocurrences, secondK)
				fmt.Println("aferDelet")
				printMapStringInt(ocurrences)
				number = first
				foundSingle = true
			} else {
				return ("NO")
			}
		}
	}

	// switch diff := first - second; diff {
	// case -1:
	// 	number = first
	// case 1:
	// 	number = second
	// case 0:
	// 	number = first
	// }

	for _, value := range ocurrences {

		if number-value == -1 || number-value == 1 && !foundSingle {
			foundSingle = true
		} else if number-value == -1 || number-value == 1 && foundSingle {
			return ("NO")
		} else if number-value < -1 || number-value > 0 {
			return ("NO")
		}
	}
	return "YES"
}

func mainSHERLOCK() {

	fmt.Println(isValid("xxxaabbccrry"))

	// aabbc

	// xxxaabbccrry
}
