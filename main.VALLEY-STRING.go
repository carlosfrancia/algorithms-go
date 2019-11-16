package main

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var level, result int32

	for _, letter := range s {
		if string(letter) == "D" {
			if level == 0 {
				result++
			}
			level--
		} else {
			level++
		}
	}
	return result
}

func main_VALLEY() {

	result := countingValleys(8, "DDUUUUDD")
	println(result)
}
