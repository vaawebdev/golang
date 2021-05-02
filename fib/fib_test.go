package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibRecErrorWhenNIsLessThanOne(t *testing.T) {
	_, err := FibRec(0)
	assert.EqualError(t, err, "n cannot be less than one")
}

func TestFibRecOne(t *testing.T) {
	res, err := FibRec(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, res)
}

func TestFibRecFive(t *testing.T) {
	res, err := FibRec(5)
	assert.Nil(t, err)
	assert.Equal(t, 5, res)
}

func TestFibRecThousand(t *testing.T) {
	res, err := FibRec(1000)
	assert.Nil(t, err)
	assert.Equal(t, 817770325994397771, res)
}

func TestFibLoopErrorWhenNIsLessThanOne(t *testing.T) {
	_, err := FibLoop(0)
	assert.EqualError(t, err, "n cannot be less than one")
}

func TestFibLoopOne(t *testing.T) {
	res, err := FibLoop(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, res)
}

func TestFibLoopFive(t *testing.T) {
	res, err := FibLoop(5)
	assert.Nil(t, err)
	assert.Equal(t, 5, res)
}

func TestFibLoopThousand(t *testing.T) {
	res, err := FibLoop(1000)
	assert.Nil(t, err)
	assert.Equal(t, 817770325994397771, res)
}
