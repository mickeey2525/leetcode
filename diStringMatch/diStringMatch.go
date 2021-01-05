package main

import (
	"fmt"
)

func diStringMatch(S string) []int {
	result := make([]int, len(S) + 1)
	max := len(S)

	min := 0

	index := 0
	for _, ch := range S{
		if ch == 'I'{
			result[index] = min
			min++
		}else{
			result[index] = max
			max--
		}
		index++
	}

	result[index] = min
	return result
}

func main() {
		a := diStringMatch("IDID")
		fmt.Println(a)
}
