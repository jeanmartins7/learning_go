package array

func linear_search(neele int, haystack []int) int {

	for i := range haystack {
		currenct_haystack := haystack[i]

		if currenct_haystack == neele {
			return i
		}
	}
	return -1
}
