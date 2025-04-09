package main

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	p := dummy
	for i := 0; i < n; i++ {
		head = head.Next
	}
	for head != nil {
		head = head.Next
		p = p.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}
