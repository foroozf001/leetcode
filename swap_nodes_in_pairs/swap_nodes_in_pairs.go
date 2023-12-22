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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	a := head
	b := head.Next

	a.Next = swapPairs(a.Next.Next)
	b.Next = a

	return b
}
