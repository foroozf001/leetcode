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

// Iterative
func (head *ListNode) reverseKListIter(k int) *ListNode {
	var prev, orig, curr *ListNode = nil, head, head
	for curr != nil && k > 0 {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		k--
	}
	orig.Next = curr
	return prev
}

func reverseKGroupIter(head *ListNode, k int) *ListNode {
	return head.reverseKListIter(k)
}

// Recursive
func (head *ListNode) reverseKGroupRecur(k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	r := head.Next.reverseKGroupRecur(k - 1)
	head.Next.Next = head
	head.Next = nil
	return r
}

func (l *ListNode) getNodeRecur(n int) *ListNode {
	if l == nil {
		return nil
	} else if n > 1 {
		return l.Next.getNodeRecur(n - 1)
	}
	return l
}

func reverseKGroupRecur(head *ListNode, k int) *ListNode {
	var orig, last *ListNode = head, head.getNodeRecur(k + 1)
	rev := head.reverseKGroupRecur(k)
	orig.Next = last
	return rev
}
