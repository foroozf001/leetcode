package main

// "reflect"

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func (l *ListNode) Equals(m *ListNode) bool {
// 	return reflect.DeepEqual(l, m)
// }

func (l *ListNode) Print() {
	print(l.Val)
	if l.Next != nil {
		l.Next.Print()
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 != nil {
		return l2
	}

	sum := l1.Val + l2.Val
	carry := sum / 10
	sum %= 10

	next := addTwoNumbers(l1.Next, l2.Next)

	if carry > 0 {
		next = addTwoNumbers(&ListNode{carry, nil}, next)
	}

	return &ListNode{sum, next}
}
