package day3

/*
It’s amusing that all the names perfectly describe what they do. We essentially traverse the array from the beginning to find the smallest value. Then, we swap it with the first element.

Similar to bubble sort, we assume that the initial element is already sorted. Therefore, in the next step, we don’t start from the first element but from the second. Once again, we attempt to find the smallest element and swap it with the element in the second position.

It is possible that we won’t find anything smaller than the starting element, and in such cases, no swaps are performed.

The algorithm
Imagine you have an array of numbers:

[5,3,8,4,2]

First Pass

Start with the assumption that the first element (5) is already the smallest.
Compare it with the second element (3). Since 3 is smaller, swap them.
Array becomes: [3,5,8,4,2]
Second Pass

Assume the first two elements are sorted.
Move to the next position (5) and compare it with the remaining elements.
Find that 4 is smaller than 5, so swap them.
Array becomes: [3,4,8,5,2]
Third Pass

Assume the first three elements are sorted.
Move to the next position (8) and compare it with the remaining elements.
Find that 5 is smaller than 8, so swap them.
Array becomes: [3,4,5,8,2]
Fourth Pass

Assume the first four elements are sorted.
Move to the next position (8) and compare it with the remaining element (2).
Find that 2 is smaller than 8, so swap them.
Array becomes: [3,4,5,2,8]
*/

func selectionSort(arr []int) []int {

	for ii := 0; ii < len(arr); ii++ {
		minIndex := ii

		for jj := 0; jj < len(arr); jj++ {
			if arr[jj] < arr[minIndex] {
				minIndex = jj
			}
		}

		arr[ii], arr[minIndex] = arr[minIndex], arr[ii]

	}

	return arr
}
