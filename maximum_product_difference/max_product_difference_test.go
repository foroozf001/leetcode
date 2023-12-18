package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
  case1 := []int{5,6,2,7,4}
  case2 := []int{4,2,5,9,7,4,8}

  assert.Equal(t, 34, MaxProductDifference(case1))
  assert.Equal(t, 64, MaxProductDifference(case2))
}
