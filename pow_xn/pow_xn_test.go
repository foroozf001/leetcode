package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Test(t *testing.T) {
	assert.Equal(t, 1024.00000, roundFloat(myPow(2.00000, 10), 4))
	assert.Equal(t, 9.26100, roundFloat(myPow(2.10000, 3), 4))
	assert.Equal(t, 0.25, roundFloat(myPow(2.00000, -2), 4))
}
