package array

func bubble_sort(list []int) []int {

	for ii := 0; ii < len(list); ii++ {
		for jj := 0; jj < len(list)-1-ii; jj++ {
			if list[jj] > list[jj+1] {
				temp := list[jj]
				list[jj] = list[jj+1]
				list[jj+1] = temp
			}
		}
	}

	return list
}
