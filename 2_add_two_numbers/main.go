package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	head := result
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
		}
		if l2 != nil {
			n2 = l2.Val
		}

		if (n1 + n2 + carry) < 10 {
			result.Next = &ListNode{Val: n1 + n2 + carry}
			carry = 0
		} else {
			result.Next = &ListNode{Val: n1 + n2 + carry - 10}
			carry = 1
		}
		result = result.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		result.Next = &ListNode{Val: carry}
	}
	return head
}
