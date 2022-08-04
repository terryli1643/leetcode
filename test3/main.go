package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	s := strconv.Itoa(root.Val) + "," + Serialize(root.Left) + "," + Serialize(root.Right)
	fmt.Println(s)
	return s
}

func Deserialize(s string) *TreeNode {
	str := strings.Split(s, ",")
	return buildTree(&str)

}

func buildTree(s *[]string) *TreeNode {
	rootValue := (*s)[0]
	*s = (*s)[1:]
	if rootValue == "#" {
		return nil
	}
	val, _ := strconv.Atoi(rootValue)
	root := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	root.Left = buildTree(s)
	root.Right = buildTree(s)
	return root
}

func main() {
	Serialize(&TreeNode{Val: 1, Right: &TreeNode{Val: 2}})
}
