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
	fmt.Println("---------------------------------------------------------")
	fmt.Println("                  SUBMENU KELOLA NEGARA                  ")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("1. Tambah Negara")
	fmt.Println("2. Edit Negara")
	fmt.Println("3. Hapus Negara")
	fmt.Println("4. Keluar")
	fmt.Println("---------------------------------------------------------")
}

func SubmenuKelolaMedali() {
	fmt.Println("---------------------------------------------------------")
	fmt.Println("                  SUBMENU KELOLA MEDALI                  ")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("1. Tambah Medali")
	fmt.Println("2. Edit Medali")
	fmt.Println("3. Hapus Medali")
	fmt.Println("4. Keluar")
	fmt.Println("---------------------------------------------------------")
}

func TampilkanRanking() {
	fmt.Println("TEST!!! untuk fungsi TampilkanRanking()")
}
