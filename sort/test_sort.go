package sort

import "fmt"

func Test_quick_sort() {

	arr := []int{2, 3, 1, 5, 6, 7, 8, 4}
	fmt.Println("Iniciando ordenação com uma lista: ", arr)
	quick_sort(arr, 0, len(arr)-1)
	fmt.Println("Quick sort: ", arr)
}

func Test_merge_sort() {
	arr := []int{2, 3, 1, 5, 6, 7, 8, 4}
	fmt.Println("Iniciando ordenação com uma lista: ", arr)
	mergeSort(arr)
	fmt.Println("Merge sort: ", arr)
}
