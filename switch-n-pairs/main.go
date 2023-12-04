package main

/*
给定一个单链表的头节点 head,实现一个调整单链表的函数，使得每K个节点之间为一组进行逆序，并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。（不能使用队列或者栈作为辅助）
例如：
链表:1->2->3->4->5->6->7->8->null, K = 3。那么 6->7->8，3->4->5，1->2各位一组。调整后：1->2->5->4->3->8->7->6->null。其中 1，2不调整，因为不够一组。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func Group(l *ListNode, k int) *ListNode {
	var h, p, cur *ListNode
	p = l
	cur = l

	for i := 0; i < k-1; i++ {
		if cur != nil {
			cur = cur.Next
		} else {
			cur = nil
		}
	}

	if cur == nil {
		return p
	} else {
		h = cur.Next
		cur.Next = nil
	}

	reverse(p, p.Next)

	g := Group(h, k)
	p.Next = g
	return cur
}

func reverse(p, l *ListNode) {
	if l.Next != nil {
		reverse(l, l.Next)
	}

	l.Next = p
}

func main() {
	// l9 := &ListNode{Val: 9, Next: nil}
	// l8 := &ListNode{Val: 8, Next: l9}
	// l7 := &ListNode{Val: 7, Next: l8}
	// l6 := &ListNode{Val: 6, Next: l7}
	// l5 := &ListNode{Val: 5, Next: l6}
	// l4 := &ListNode{Val: 4, Next: l5}
	// l3 := &ListNode{Val: 3, Next: l4}
	// l2 := &ListNode{Val: 2, Next: l3}
	// l1 := &ListNode{Val: 1, Next: l2}
	// h := Group(l1, 3)
	// for h != nil {
	// 	fmt.Print(h.Val)
	// 	h = h.Next
	// }
	var c chan string
	c <- ""
}
