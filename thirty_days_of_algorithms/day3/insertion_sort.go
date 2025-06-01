package day3

/*
For me, insertion sort was the most challenging. It’s a combination of both bubble and selection sort. We start from the second position and compare the elements to the left of this position.

Whenever an element is larger than the starting one, we shift the element one index to the right. This process continues until we reach the beginning of the array or encounter an element that isn’t larger than the starting one.

In the next step, we start from the third position and compare elements to the left and so on.

The algorithm
Imagine you have an array of numbers:

[5,3,8,4,2]

First Pass

Start with the assumption that the first element (5) is already sorted.
Move to the second position (3) and compare it with the sorted elements.
Since 3 is smaller than 5, shift 5 to the right.
Array becomes: [3,5,8,4,2]
Second Pass

Move to the third position (8) and compare it with the sorted elements.
Since 8 is greater than 5, leave it in its place.
Array remains: [3,5,8,4,2]
Third Pass

Move to the fourth position (4) and compare it with the sorted elements.
Since 4 is smaller than 8, shift 8 to the right and 5 to the right.
Array becomes: [3,4,5,8,2]
Fourth Pass

Move to the fifth position (2) and compare it with the sorted elements.
Since 2 is smaller than 8, shift 8 to the right, 5 to the right, and 4 to the right.
Array becomes: [2,3,4,5,8]
*/

func insertionSort(arr []int) []int {

	for ii := 0; ii < len(arr); ii++ {

		value := arr[ii]
		jj := ii - 1

		for jj >= 0 && value < arr[jj] {
			arr[jj+1] = arr[jj]
			jj = jj - 1
		}

		arr[jj+1] = value
	}

	return arr
}
