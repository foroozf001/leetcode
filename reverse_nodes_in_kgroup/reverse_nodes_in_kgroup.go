package main

func main() {

}

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

func length(l *ListNode) int {
	if l == nil {
		return 0
	}
	length := 1
	for l.Next != nil {
		l = l.Next
		length++
	}
	return length
}

// Iterative
func reverseKGroupIter(head *ListNode, k int) *ListNode {
	i := k
	if length(head) < i {
		return head
	}
	var prev, curr *ListNode = nil, head
	for curr != nil && i > 0 {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		i--
	}
	head.Next = reverseKGroupIter(curr, k)
	return prev
}

// Recursive
func reverseKListRecur(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	reverse := reverseKListRecur(head.Next, k-1)
	head.Next.Next = head
	head.Next = nil
	return reverse
}

func getNode(l *ListNode, n int) *ListNode {
	if l == nil {
		return nil
	} else if n > 1 {
		return getNode(l.Next, n-1)
	}
	return l
}

func reverseKGroupRecur(head *ListNode, k int) *ListNode {
	if length(head) < k {
		return head
	}
	var next *ListNode = getNode(head, k+1)
	reverse := reverseKListRecur(head, k)
	head.Next = reverseKGroupRecur(next, k)
	return reverse
}
