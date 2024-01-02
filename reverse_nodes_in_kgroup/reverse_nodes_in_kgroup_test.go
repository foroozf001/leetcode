package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	l1e := &ListNode{5, nil}
	l1d := &ListNode{4, l1e}
	l1c := &ListNode{3, l1d}
	l1b := &ListNode{2, l1c}
	l1a := &ListNode{1, l1b}

	l2e := &ListNode{5, nil}
	l2d := &ListNode{3, l2e}
	l2c := &ListNode{4, l2d}
	l2b := &ListNode{1, l2c}
	l2a := &ListNode{2, l2b}

	assert.Equal(t, l2a, reverseKGroupRecur(l1a, 2))

	l1e = &ListNode{5, nil}
	l1d = &ListNode{4, l1e}
	l1c = &ListNode{3, l1d}
	l1b = &ListNode{2, l1c}
	l1a = &ListNode{1, l1b}

	assert.Equal(t, l2a, reverseKGroupIter(l1a, 2))

	l1e = &ListNode{5, nil}
	l1d = &ListNode{4, l1e}
	l1c = &ListNode{3, l1d}
	l1b = &ListNode{2, l1c}
	l1a = &ListNode{1, l1b}

	l2e = &ListNode{5, nil}
	l2d = &ListNode{4, l2e}
	l2c = &ListNode{1, l2d}
	l2b = &ListNode{2, l2c}
	l2a = &ListNode{3, l2b}

	assert.Equal(t, l2a, reverseKGroupRecur(l1a, 3))

	l1e = &ListNode{5, nil}
	l1d = &ListNode{4, l1e}
	l1c = &ListNode{3, l1d}
	l1b = &ListNode{2, l1c}
	l1a = &ListNode{1, l1b}

	assert.Equal(t, l2a, reverseKGroupIter(l1a, 3))
}
