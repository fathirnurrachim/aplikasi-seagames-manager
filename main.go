package main

import (
	"fmt"
)

func main() {
	var data Data
	var nData, pilih, pilihSubmenu int

	for {
		Menu()
		fmt.Println("Pilih (1/2/3/4)?")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			for {
				SubmenuKelolaNegara()
				fmt.Println("Pilih (1/2/3/4)?")
				fmt.Scan(&pilihSubmenu)

				switch pilihSubmenu {
				case 1:
					TambahNegara(data, nData)
				case 2:
					EditNegara(data, nData)
				case 3:
					HapusNegara(data, nData)
				}

				if pilihSubmenu == 4 {
					return
				}
			}

		case 2:
			for {
				SubmenuKelolaMedali()
				fmt.Println("Pilih (1/2/3/4)?")
				fmt.Scan(&pilihSubmenu)

				switch pilihSubmenu {
				case 1:
					TambahMedali(data, nData)
				case 2:
					EditMedali(data, nData)
				case 3:
					HapusMedali(data, nData)
				}

				if pilihSubmenu == 4 {
					return
				}
			}

		case 3:
			TampilkanRanking()
		}

		if pilih == 4 {
			return
		}
	}
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
