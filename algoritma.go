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

func BinarySearch(data Data, nData, idCari int) int {
	var index, left, right, mid int

	index = -1
	left = 0
	right = nData - 1

	for left <= right && index == -1 {
		mid = (left + right) / 2

		if data[mid].id == idCari {
			index = mid
		} else if data[mid].id < idCari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return index
}

func SelectionSort() {
	fmt.Println("TEST!!! untuk fungsi SelectionSort()")
}

func InsertionSort() {
	fmt.Println("TEST!!! untuk fungsi InsertionSort()")
}
