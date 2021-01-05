package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(a, b int) bool { return boxTypes[a][1] > boxTypes[b][1] })
	ans := 0
	for _, box := range boxTypes {
		fmt.Println(box)
		if box[0] <= truckSize {
			ans += box[0] * box[1]
			truckSize -= box[0]
		} else {
			ans += box[1] * truckSize
			break
		}
	}
	return ans
}

func main(){
	sample := [][]int{{1,3},{2,2},{3, 1}}
	a := maximumUnits(sample, 4)
	fmt.Println(a)
}
