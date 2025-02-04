package main

// LCR 024. 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return reverse(nil, head)
}

func reverse(parent *ListNode, child *ListNode) *ListNode {
	next := child.Next
	child.Next = parent
	if next == nil {
		return child
	}
	return reverse(child, next)
}
