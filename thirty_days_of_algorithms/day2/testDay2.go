package day2

import "fmt"

func StartTestDay2() {

	arr := make([]int, 5)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	result := simpleSearch(1, arr)
	fmt.Println(result)

}
