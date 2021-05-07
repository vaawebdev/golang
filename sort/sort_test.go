package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func makeRange(len int) []int {
	arr := make([]int, len)

	for i := range arr {
		arr[i] = i + 1
	}

	dest := make([]int, len)
	perm := rand.Perm(len)

	for i, v := range perm {
		dest[v] = arr[i]
	}

	return dest
}

func TestBubbleSortEmptyArray(t *testing.T) {
	arr := make([]int, 0)
	BubbleSort(arr, func(a int, b int) int { return a - b })
	assert.Empty(t, arr)
}

func TestBubbleSort(t *testing.T) {
	arr := makeRange(3)
	BubbleSort(arr, func(a int, b int) int { return a - b })
	assert.Equal(t, []int{1, 2, 3}, arr)
}

func TestBubbleSortReverse(t *testing.T) {
	arr := makeRange(3)
	BubbleSort(arr, func(a int, b int) int { return b - a })
	assert.Equal(t, []int{3, 2, 1}, arr)
}
