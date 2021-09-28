package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum, anchor := 0, new(ListNode)
	for node := anchor; l1 != nil || l2 != nil || sum > 0; node = node.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{sum % 10, nil}
		sum /= 10
	}
	return anchor.Next
}

func main() {
	l1, l2 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	addTwoNumbers(l1, l2)
	fmt.Println("done")
}
