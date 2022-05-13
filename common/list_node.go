package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(nums []int) *ListNode {
	head := ListNode{}
	p := &head
	for _, n := range nums {
		p.Next = &ListNode{n, nil}
		p = p.Next
	}
	return head.Next
}
