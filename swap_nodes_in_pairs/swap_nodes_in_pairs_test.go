package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	l1d := &ListNode{4, nil}
	l1c := &ListNode{3, l1d}
	l1b := &ListNode{2, l1c}
	l1a := &ListNode{1, l1b}

	l2d := &ListNode{3, nil}
	l2c := &ListNode{4, l2d}
	l2b := &ListNode{1, l2c}
	l2a := &ListNode{2, l2b}

	assert.Equal(t, l2a, swapPairs(l1a))

	l1a = nil
	l2a = nil

	assert.Equal(t, l2a, swapPairs(l1a))

	l1a = &ListNode{0, nil}
	l2a = &ListNode{0, nil}

	assert.Equal(t, l2a, swapPairs(l1a))
}
