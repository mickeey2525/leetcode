package main

import "fmt"

type Node struct {
	Val int
	Children []*Node
}

func postorder(root *Node) []int {
	res := []int{}
	dfs(root, &res)
	return res
}

func dfs(root *Node, res *[]int) {
	if root == nil { return }
	for _, child := range root.Children { dfs(child, res) }
	*res = append(*res, root.Val)
}

func main() {
	a := []int{1,nil,3,2,4,nil,5,6}
	ans := postorder(a)
	fmt.Println(ans)
}
