package array

func selection_sort(list []int) []int {

	size := len(list)

	for ii := 0; ii < size; ii++ {
		min_index := ii

		for jj := 0; jj < (ii+1)-size; jj++ {
			if list[jj] < list[min_index] {
				min_index = jj
			}
		}

		temp := list[ii]
		list[ii] = list[min_index]
		list[min_index] = temp
	}

	return list
}
