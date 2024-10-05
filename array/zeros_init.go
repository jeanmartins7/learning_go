package array

func zeros_init(arr []int) []int {
	j := len(arr) - 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			arr[j], arr[i] = arr[i], arr[j]
			j--
		}
	}

	return arr
}
