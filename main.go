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
	fmt.Println("------------------------------")
	fmt.Println("             MENU             ")
	fmt.Println("------------------------------")
	fmt.Println("1. Kelola Negara")
	fmt.Println("2. Kelola Medali")
	fmt.Println("3. Tampilkan Ranking")
	fmt.Println("4. Keluar")
	fmt.Println("------------------------------")
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
