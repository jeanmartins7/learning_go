package array

import (
	"fmt"
)

func TestLinearSearch() {

	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	resultLiearSearch := linear_search(4, list1)

	fmt.Print("O resultado para o n: ", 4, " na posicao = ", resultLiearSearch)
}

func TestBinarySearch() {

	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultLiearSearch := binary_search(3, list1)

	fmt.Print("O resultado para o n: ", 3, " na posicao = ", resultLiearSearch)

}

func TestBubbleSort() {
	list := []int{3, 5, 2, 8, 1}
	sorted := bubble_sort(list)

	fmt.Println(sorted)
}

func TestInsertionSort() {
	list := []int{3, 5, 2, 8, 1}
	sorted := insertion_sort(list)

	fmt.Println(sorted)
}
