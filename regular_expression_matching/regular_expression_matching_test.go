package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, false, IsMatch("aa", "a"))
	assert.Equal(t, true, IsMatch("aa", "a*"))
	assert.Equal(t, true, IsMatch("ab", ".*"))
}
