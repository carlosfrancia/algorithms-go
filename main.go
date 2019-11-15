package main

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	var result int32
	socks := make(map[int32]int32)

	for _, value := range ar {
		if socks[value] == 0 {
			println("EMPTY")
			socks[value] = 1
		} else {
			println("NOT EMPTY")
			socks[value] = socks[value] + 1
		}
	}
	for _, value := range socks {
		result = result + value/2

	}
	return result
}

func main() {

	result := sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})

	println(result)
}
