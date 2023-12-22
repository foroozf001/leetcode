package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	l1c := &ListNode{4, nil}
	l1b := &ListNode{2, l1c}
	l1a := &ListNode{1, l1b}

	l2c := &ListNode{4, nil}
	l2b := &ListNode{3, l2c}
	l2a := &ListNode{1, l2b}

	l3f := &ListNode{4, nil}
	l3e := &ListNode{4, l3f}
	l3d := &ListNode{3, l3e}
	l3c := &ListNode{2, l3d}
	l3b := &ListNode{1, l3c}
	l3a := &ListNode{1, l3b}

	assert.Equal(t, l3a, mergeTwoLists(l1a, l2a))

	l1a = nil
	l2a = nil
	l3a = nil

	assert.Equal(t, l3a, mergeTwoLists(l1a, l2a))

	l1a = nil
	l2a = &ListNode{0, nil}
	l3a = &ListNode{0, nil}

	assert.Equal(t, l3a, mergeTwoLists(l1a, l2a))
}
