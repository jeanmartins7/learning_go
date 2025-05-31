package day2

func simpleSearch(needle int, arr []int) int {

	heystack := 0

	for ii := 0; ii < len(arr); ii++ {

		heystack = arr[ii]
		if heystack == needle {
			return ii
		}

	}

	return -1
}
