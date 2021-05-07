package bsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeRange(min int, max int) []int {
	arr := make([]int, max-min+1)

	for i := range arr {
		arr[i] = min + i
	}

	return arr
}

func TestBSearchEmptyArray(t *testing.T) {
	res := Bsearch(1, []int{})
	assert.Equal(t, -1, res)
}

func TestBSearchNotFound(t *testing.T) {
	res := Bsearch(0, []int{1, 2, 3})
	assert.Equal(t, -1, res)
}

func TestBSearchFirst(t *testing.T) {
	res := Bsearch(1, []int{1, 2, 3})
	assert.Equal(t, 0, res)
}

func TestBSearchLast(t *testing.T) {
	res := Bsearch(9999999999999, []int{-80, 792, 1290, 88198259, 9999999999999})
	assert.Equal(t, 4, res)
}

func TestBSearchBig(t *testing.T) {
	res := Bsearch(19999999, makeRange(500, 99999999))
	assert.Equal(t, 19999499, res)
}
