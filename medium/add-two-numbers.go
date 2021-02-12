package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l ListNode
	remember := 0
	lx := &l

	for l1 != nil || l2 != nil || remember != 0 {
		v1 := 0
		v2 := 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		v := v1 + v2 + remember
		remember = v / 10
		v = v % 10
		lx.Val = v

		if l1 != nil || l2 != nil || remember != 0 {
			var newNode ListNode
			lx.Next = &newNode
			lx = lx.Next
		}
	}

	return &l
}

func main() {
	var l1 ListNode
	var l2 ListNode

	l1.Val = 5
	l2.Val = 5

	fmt.Println(addTwoNumbers(&l1, &l2))
}
