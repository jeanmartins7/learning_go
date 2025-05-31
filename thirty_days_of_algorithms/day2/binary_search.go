package day2

import "sort"

func binary_search(needle int, arr []int) int {

	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == needle {
			return mid
		}

		if arr[mid] < needle {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1

}
