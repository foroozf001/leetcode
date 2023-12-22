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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 != nil && list2 == nil {
		return list1
	} else if list1 == nil && list2 != nil {
		return list2
	}

	if list1.Val <= list2.Val {
		return &ListNode{list1.Val, mergeTwoLists(list1.Next, list2)}
	}

	return &ListNode{list2.Val, mergeTwoLists(list1, list2.Next)}
}
