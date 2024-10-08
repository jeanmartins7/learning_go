package sort

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quick_sort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quick_sort(arr, low, pi-1)
		quick_sort(arr, pi+1, high)
	}
}
