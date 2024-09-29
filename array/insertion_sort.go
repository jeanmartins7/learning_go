package array

func insertion_sort(list []int) []int {

	for ii := 0; ii < len(list); ii++ {
		value := list[ii]
		jj := ii - 1

		for jj >= 0 && value < list[jj] {
			list[jj+1] = list[jj]
			jj = jj - 1
		}
		list[jj+1] = value
	}

	return list
}
