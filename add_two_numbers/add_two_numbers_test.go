package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	l1c := &ListNode{3, nil}
	l1b := &ListNode{4, l1c}
	l1a := &ListNode{2, l1b}

	l2c := &ListNode{4, nil}
	l2b := &ListNode{6, l2c}
	l2a := &ListNode{5, l2b}

	l3c := &ListNode{8, nil}
	l3b := &ListNode{0, l3c}
	l3a := &ListNode{7, l3b}

	assert.Equal(t, l3a, AddTwoNumbers(l1a, l2a))

	l1a = &ListNode{0, nil}
	l2a = &ListNode{0, nil}
	l3a = &ListNode{0, nil}

	assert.Equal(t, l3a, AddTwoNumbers(l1a, l2a))

	l1g := &ListNode{9, nil}
	l1f := &ListNode{9, l1g}
	l1e := &ListNode{9, l1f}
	l1d := &ListNode{9, l1e}
	l1c = &ListNode{9, l1d}
	l1b = &ListNode{9, l1c}
	l1a = &ListNode{9, l1b}

	l2d := &ListNode{9, nil}
	l2c = &ListNode{9, l2d}
	l2b = &ListNode{9, l2c}
	l2a = &ListNode{9, l2b}

	l3h := &ListNode{1, nil}
	l3g := &ListNode{0, l3h}
	l3f := &ListNode{0, l3g}
	l3e := &ListNode{0, l3f}
	l3d := &ListNode{9, l3e}
	l3c = &ListNode{9, l3d}
	l3b = &ListNode{9, l3c}
	l3a = &ListNode{8, l3b}

	assert.Equal(t, l3a, AddTwoNumbers(l1a, l2a))
}
