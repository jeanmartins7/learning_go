package sort

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i = j + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func quick_sort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quick_sort(arr, low, pi-1)
		quick_sort(arr, pi+1, high)
	}
}
