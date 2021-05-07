package sort

func BubbleSort(arr []int, matcher func(a int, b int) int) {
	len := len(arr)

	if len == 1 {
		return
	}

	swapped := true
	counter := 0
	var tmp int

	for swapped {
		swapped = false
		counter += 1

		for i := 0; i < len-counter; i++ {
			if matcher(arr[i], arr[i+1]) > 0 {
				tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
				swapped = true
			}
		}
	}
}
