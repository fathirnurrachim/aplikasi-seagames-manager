package main

import (
	"fmt"
	"strings"
)

var negara = map[string]map[string]int{
	"INA": {"emas": 0, "perak": 0, "perunggu": 0},
	"MAS": {"emas": 0, "perak": 0, "perunggu": 0},
	"SGP": {"emas": 0, "perak": 0, "perunggu": 0},
}

func KelolaMedali() {
	var pilih int
	selesai := false

	for !selesai {
		fmt.Println("\n========== MENU MEDALI ==========")
		fmt.Println("1. Tambah Medali")
		fmt.Println("2. Edit Medali")
		fmt.Println("3. Hapus Medali")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih Menu: ")

		fmt.Scanln(&pilih)

		if pilih == 4 {
			selesai = true
		} else if pilih >= 1 && pilih <= 3 {
			var kode string
			fmt.Print("Masukkan Kode Negara (INA/MAS/SGP): ")
			fmt.Scanln(&kode)

			kode = strings.ToUpper(kode)

			if _, ada := negara[kode]; ada {
				var jenis string
				fmt.Print("Jenis Medali (emas/perak/perunggu): ")
				fmt.Scanln(&jenis)

				jenis = strings.ToLower(jenis)

				if jenis == "emas" || jenis == "perak" || jenis == "perunggu" {
					if pilih == 1 {
						var jumlah int
						fmt.Print("Jumlah Medali: ")
						fmt.Scanln(&jumlah)

						negara[kode][jenis] += jumlah
						fmt.Println("Medali berhasil ditambahkan")

					} else if pilih == 2 {
						var jumlahBaru int
						fmt.Print("Jumlah Baru: ")
						fmt.Scanln(&jumlahBaru)

						negara[kode][jenis] = jumlahBaru
						fmt.Println("Medali berhasil diupdate")

					} else if pilih == 3 {
						negara[kode][jenis] = 0
						fmt.Println("Medali berhasil dihapus")
					}

					fmt.Printf("\n%s -> Emas:%d Perak:%d Perunggu:%d\n",
						kode,
						negara[kode]["emas"],
						negara[kode]["perak"],
						negara[kode]["perunggu"])

				} else {
					fmt.Println("Jenis medali tidak valid")
				}
			} else {
				fmt.Println("Negara tidak ditemukan")
			}
		} else {
			fmt.Println("Menu tidak valid! Silakan pilih 1-4.")
		}
	}
}
