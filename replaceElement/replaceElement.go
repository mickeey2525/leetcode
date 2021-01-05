package main

import (
	"fmt"
)

func replaceElements(arr []int) []int {
	var ans []int
	if len(arr) == 1 {
		ans = append(ans, -1)
	}
	for i,_ := range arr {
		if i+1 == len(arr){
			break
		}
		ans = append(ans, maxInt(arr[i+1:]))
	}

	ans = append(ans,-1)

	return ans
}

func maxInt(a []int) int {
	max := 0
	for i:=0; i< len(a); i++ {
		if max < a[i]{
			max = a[i]
		}
	}
	return max
}

func main(){
	a := []int{17,18,5,4,6,1}
	b := replaceElements(a)
	fmt.Println(b)
}
