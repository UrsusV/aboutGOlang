package main

import "fmt"

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	return arr
}

func main() {
	arr := []int{5, 2, 1, 3, 4, 0}
	var sortedArr []int
	sortedArr = insertionSort(arr)
	fmt.Println(sortedArr)
}
