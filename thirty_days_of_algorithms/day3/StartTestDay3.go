package day3

import (
	"fmt"
	"strconv"
)

//[5,3,8,4,2]

func StartTestDay3() {

	arr := make([]int, 5)
	arr[0] = 5
	arr[1] = 3
	arr[2] = 8
	arr[3] = 4
	arr[4] = 2

	result := bubbleSort(arr)

	fmt.Printf("Result Bubble sort: ")
	for ii := 0; ii < len(result); ii++ {
		fmt.Printf(strconv.Itoa(result[ii]))
	}

	resetArr(arr)

	result = selectionSort(arr)

	fmt.Printf("Result Selection sort: ")
	for ii := 0; ii < len(result); ii++ {
		fmt.Printf(strconv.Itoa(result[ii]))
	}

	resetArr(arr)

	result = insertionSort(arr)

	fmt.Printf("Result Insertion sort: ")
	for ii := 0; ii < len(result); ii++ {
		fmt.Printf(strconv.Itoa(result[ii]))
	}
}

func resetArr(arr []int) []int {
	arr[0] = 5
	arr[1] = 3
	arr[2] = 8
	arr[3] = 4
	arr[4] = 2

	fmt.Print("\n")

	fmt.Printf("reset array: ")
	for ii := 0; ii < len(arr); ii++ {
		fmt.Printf(strconv.Itoa(arr[ii]))
	}

	fmt.Print("\n")

	return arr
}
