package main

func countCharacters(words []string, chars string) int {
	tmp := make(map[rune]int)
	length := 0

	for _, v := range chars {
		tmp[v]++
	}
	for _, v := range words {
		temp := make(map[rune]int)
		for k, v := range tmp {
			temp[k] = v
		}
		for I, V := range v {
			if _, ok := temp[V]; !ok {
				break
			}
			if _, ok := temp[V]; ok {
				temp[V]--
			}
			if temp[V] < 0 {
				break
			}
			if I == len(v)-1 {
				length += len(v)
			}
		}
	}
	return length
}


func main()  {

}
