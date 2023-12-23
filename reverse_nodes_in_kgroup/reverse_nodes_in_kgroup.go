package main

func main() {
	l1e := &ListNode{5, nil}
	l1d := &ListNode{4, l1e}
	l1c := &ListNode{3, l1d}
	l1b := &ListNode{2, l1c}
	l1a := &ListNode{1, l1b}

	l1a.GetNodeIteratively(5).Print()
	l1a.GetNodeRecursively(5).Print()
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

func (l *ListNode) GetNodeRecursively(k int) *ListNode {
	if k > 0 && l.Next != nil {
		return l.Next.GetNodeRecursively(k - 1)
	}
	return l
}

func (l *ListNode) GetNodeIteratively(k int) *ListNode {
	var j int
	for ; j < k && l.Next != nil; j++ {
		l = l.Next
	}
	return l
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}
