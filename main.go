package main

import (
	"fmt"
)

func main() {
	fmt.Println("TEST!!! untuk fungsi main()")

	Menu()
	SubmenuKelolaNegara()
	SubmenuKelolaMedali()
	TampilkanRanking()

	KelolaNegara()
	KelolaMedali()

	SequentialSearch()
	BinarySearch()
	SelectionSort()
	InsertionSort()
}

func Menu() {
	fmt.Println("TEST!!! untuk fungsi Menu()")
}

func SubmenuKelolaNegara() {
	fmt.Println("TEST!!! untuk fungsi SubmenuKelolaNegara()")
}

func SubmenuKelolaMedali() {
	fmt.Println("TEST!!! untuk fungsi SubmenuKelolaMedali()")
}

func TampilkanRanking() {
	fmt.Println("TEST!!! untuk fungsi TampilkanRanking()")
}
