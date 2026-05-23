package main

import (
	"fmt"
)

func SequentialSearch(data Data, nData, idCari int) int {
	var index, i int

	index = -1

	i = 0
	for i < nData && index == -1 {
		if data[i].id == idCari {
			index = i
		}
	}
	i++

	return index
}

func BinarySearch() {
	fmt.Println("TEST!!! untuk fungsi BinarySearch()")
}

func SelectionSort() {
	fmt.Println("TEST!!! untuk fungsi SelectionSort()")
}

func InsertionSort() {
	fmt.Println("TEST!!! untuk fungsi InsertionSort()")
}
