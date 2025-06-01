package day3

/*
Theory

Imagine you have an array of numbers:

[5,3,8,4,2]

First Pass

Compare the first two elements (5 and 3).
Since 5 is greater than 3, swap them.
Array becomes: [3,5,8,4,2]
Move to the next pair (5 and 8). No swap is needed.
Move to the next pair (8 and 4). Swap them.
Array becomes: [3,5,4,8,2]
Continue this process until the end of the array.
Second Pass

Same as the first pass but skip the last element since it’s assumed to be sorted.
Array becomes: [3,4,5,2,8]
Third Pass

Skip the last two elements.
Array becomes: [3,4,2,5,8]
Fourth Pass

Skip the last three elements.
Array becomes: [3,2,4,5,8]
Fifth Pass

Skip the last four elements.
Array becomes: [2,3,4,5,8]

*/

func bubbleSort(arr []int) []int {

	for ii := 0; ii < len(arr); ii++ {
		for jj := 0; jj < len(arr)-1; jj++ {
			if arr[jj] > arr[jj+1] {
				arr[jj], arr[jj+1] = arr[jj+1], arr[jj]
			}
		}
	}

	return arr
}
