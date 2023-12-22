package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	print(l.Val)
	if l.Next != nil {
		l.Next.Print()
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}
