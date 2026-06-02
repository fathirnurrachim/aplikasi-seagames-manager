package main

import (
	"fmt"
)

func main() {
	var data Data
	var nData, pilih, pilihSubmenu int
	var idCari, idBaru int
	var namaBaru string

	for {
		Menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			for {
				SubmenuKelolaNegara()
				fmt.Print("Pilih (1/2/3/4)? ")
				fmt.Scan(&pilihSubmenu)

				switch pilihSubmenu {
				case 1:
					fmt.Print("Masukkan Jumlah Negara: ")
					fmt.Scan(&nData)

					BacaNegara(&data, &nData)

					TambahNegara(&data, &nData, idBaru, namaBaru)

					TampilkanNegara(data, nData)
				case 2:
					// EditNegara(data, nData)
				case 3:
					// HapusNegara(data, nData)
				}

				if pilihSubmenu == 4 {
					break
				}
			}

		case 2:
			for {
				SubmenuKelolaMedali()
				fmt.Print("Pilih (1/2/3/4)? ")
				fmt.Scan(&pilihSubmenu)

				switch pilihSubmenu {
				case 1:
					// TambahMedali(data, nData)
				case 2:
					// EditMedali(data, nData)
				case 3:
					// HapusMedali(data, nData)
				}

				if pilihSubmenu == 4 {
					break
				}
			}

		case 3:
			SelectionSortDesc(data, nData)

			TampilkanNegara(data, nData)
		}

		if pilih == 4 {
			break
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

// func TampilkanRanking(data Data, nData int) {

// 	SelectionSortDesc(data, nData)

// 	fmt.Printf("%-5s %-25s %-10s %-10s %-10s\n",
// 		"ID", "Nama Negara", "Emas", "Perak", "Perunggu")

// 	for i := 0; i < nData; i++ {
// 		fmt.Printf("%-5d %-25s %-10d %-10d %-10d\n",
// 			data[i].id, data[i].nama, data[i].medali.emas, data[i].medali.perak, data[i].medali.perunggu)
// 	}
// }
