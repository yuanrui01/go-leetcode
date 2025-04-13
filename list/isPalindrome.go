package main

// LCR 027. 回文链表
func isPalindrome(head *ListNode) bool {
	middle := findMiddle(head)
	head2 := reverse(middle)
	for head2 != nil {
		if head2.Val != head.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
