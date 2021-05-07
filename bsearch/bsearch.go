package bsearch

import (
	"math"
)

func Bsearch(needle int, arr []int) int {
	len := len(arr)

	if len == 0 {
		return -1
	}

	if len == 1 {
		if arr[0] == needle {
			return 0
		}

		return -1
	}

	start := 0
	end := len - 1
	center := int(math.Floor(float64(end-start) / 2))

	for true {
		if arr[center] == needle {
			return center
		} else if center == 0 || center == len-1 {
			return -1
		} else if arr[center] > needle {
			end = center
			center = int(math.Floor(float64(end-start)/2)) + start
		} else if arr[center] < needle {
			start = center
			center = int(math.Ceil(float64(end-start)/2)) + start
		}

	}

	return -1
}
