package main

import "strings"

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	var result int64

	if !strings.ContainsAny(s, "a") {
		return result
	}
	if int64(len(s)) >= n {

		return int64(strings.Count(s[0:n], "a"))
	}
	div := n / int64(len(s))
	mod := n % int64(len(s))
	modStr := s[0:mod]

	return int64((div * int64(strings.Count(s, "a"))) + int64(strings.Count(modStr, "a")))
}

func main_REPEATED_CHAR_IN_STRING() {

	result := repeatedString("ababa", 3)

	println(result)
}
