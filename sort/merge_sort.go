package sort

func mergeSort(array []int) {
	if len(array) > 1 {
		mid := len(array) / 2
		leftArr := make([]int, mid)
		rightArr := make([]int, len(array)-mid)

		copy(leftArr, array[:mid])
		copy(rightArr, array[mid:])

		mergeSort(leftArr)
		mergeSort(rightArr)

		i, j, k := 0, 0, 0

		for i < len(leftArr) && j < len(rightArr) {
			if leftArr[i] < rightArr[j] {
				array[k] = leftArr[i]
				i++
			} else {
				array[k] = rightArr[j]
				j++
			}
			k++
		}

		for i < len(leftArr) {
			array[k] = leftArr[i]
			i++
			k++
		}

		for j < len(rightArr) {
			array[k] = rightArr[j]
			j++
			k++
		}
	}
}
