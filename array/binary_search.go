package array

import "fmt"

func binary_search(needle int, haystack []int) int {
	low := 0
	high := len(haystack) - 1

	for low <= high {

		mid := low + (high-low)/2
		fmt.Print(mid)

		if haystack[mid] == needle {
			return mid
		} else if haystack[mid] < needle {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
