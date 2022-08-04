package main

import (
	"container/heap"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param lists ListNode类一维数组
 * @return ListNode类
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// write code here

	h := &NodeHeap{}
	heap.Init(h)
	for _, n := range lists {
		heap.Push(h, n)
	}
	mergedNode := &ListNode{}
	merge(h, mergedNode)
	for mergedNode != nil {
		log.Println(mergedNode.Val)
		mergedNode = mergedNode.Next
	}
	return nil
}

func merge(h *NodeHeap, mergedNode *ListNode) {
	if len(*h) == 0 {
		return
	}
	node := heap.Pop(h)
	if node.(*ListNode).Next != nil {
		heap.Push(h, node.(*ListNode).Next)
	}
	mergedNode.Next = node.(*ListNode)
	merge(h, mergedNode.Next)

}

type NodeHeap []*ListNode

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val <= h[j].Val
}

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	if n == 0 {
		return nil
	}
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *NodeHeap) Push(e interface{}) {
	*h = append(*h, e.(*ListNode))
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	list2 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}
	mergeKLists([]*ListNode{list1, list2})
}
