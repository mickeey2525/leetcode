package main

import (
	"fmt"
	"strconv"
)
// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/submissions/

//the maximum number of deci numbers would be equal to the maximum number in the string. Rest all can be adjusted to zero-
//987654321 =
//111111111 +
//111111110 +
//111111100 +
//111111000 +
//and so on till the maximum number in the string which is 9.

func minPartitions(n string) int {

	max := 0
	for _, num := range n {
		converted, _ := strconv.Atoi(string(num))
		if max < converted {
			max = converted
		}
	}
	return max
}

func main(){
	m := minPartitions("32")
	fmt.Println(m)
}
