package main

import (
	"fmt"
	// "strings"
)

func TambahNegara(data *Data, nData *int) {
	var index, jumlah, idBaru int
	var namaBaru string

	fmt.Print("Masukkan Jumlah Negara: ")
	fmt.Scan(&jumlah)
	fmt.Println()

	if jumlah > NMAX {
		jumlah = NMAX
	}

	for i := 0; i < jumlah; i++ {
		fmt.Print("ID Negara Baru: ")
		fmt.Scan(&idBaru)

		fmt.Print("Nama Negara Baru: ")
		fmt.Scan(&namaBaru)
		fmt.Println()

		index = SequentialSearch(*data, *nData, idBaru)

		if index == -1 {
			data[*nData].id = idBaru
			data[*nData].nama = namaBaru

			data[*nData].medali.emas = 0
			data[*nData].medali.perak = 0
			data[*nData].medali.perunggu = 0

			*nData++

			fmt.Println("NEGARA BERHASIL DITAMBAHKAN!")
			fmt.Println()
		} else {
			fmt.Println("NEGARA SUDAH ADA!")
			fmt.Println()
		}
	}
}

func EditNegara(data *Data, nData int) {
	var index, idCari int
	var namaBaru string

	fmt.Print("ID Negara Lama: ")
	fmt.Scan(&idCari)
	fmt.Println()

	index = SequentialSearch(*data, nData, idCari)

	if index != -1 {
		fmt.Print("Nama Negara Baru: ")
		fmt.Scan(&namaBaru)
		fmt.Println()

		data[index].nama = namaBaru

		fmt.Println("NEGARA BERHASIL DIEDIT!")
		fmt.Println()
	} else {
		fmt.Println("NEGARA TIDAK DITEMUKAN")
		fmt.Println()
	}
}

func HapusNegara(data *Data, nData *int) {
	var index, idCari int

	fmt.Print("ID Negara Lama: ")
	fmt.Scan(&idCari)
	fmt.Println()

	InsertionSortAsc(*data, *nData)
	index = BinarySearch(*data, *nData, idCari)

	if index != -1 {
		for i := index; i < *nData-1; i++ {
			data[i] = data[i+1]
		}

		*nData--

		fmt.Println("NEGARA BERHASIL DIHAPUS!")
		fmt.Println()
	} else {
		fmt.Println("NEGARA TIDAK DITEMUKAN!")
		fmt.Println()
	}
}

func TampilkanNegara(data Data, nData int) {

	fmt.Printf("%-5s %-25s %-10s %-10s %-10s\n",
		"ID", "Nama Negara", "Emas", "Perak", "Perunggu")

	for i := 0; i < nData; i++ {
		fmt.Printf("%-5d %-25s %-10d %-10d %-10d\n",
			data[i].id, data[i].nama, data[i].medali.emas, data[i].medali.perak, data[i].medali.perunggu)
	}

	fmt.Println()
}

// // Data negara
// var negara = map[string]string{
// 	"INA": "Indonesia",
// 	"MAS": "Malaysia",
// 	"SGP": "Singapura",
// 	"THA": "Thailand",
// 	"PHI": "Filipina",
// }

// func KelolaNegara() {
// 	for {
// 		fmt.Println("\n========== MENU NEGARA ==========")
// 		fmt.Println("1. Tambah negara")
// 		fmt.Println("2. Edit negara")
// 		fmt.Println("3. Hapus negara")
// 		fmt.Println("4. Lihat semua negara")
// 		fmt.Println("5. Kembali")
// 		fmt.Print("Pilih menu (1-5): ")

// 		var pilih int
// 		fmt.Scanln(&pilih)

// 		if pilih == 5 {
// 			fmt.Println("Kembali ke menu utama...")
// 			break
// 		}

// 		// Menu lihat semua negara
// 		if pilih == 4 {
// 			fmt.Println("\n========== DATA SEMUA NEGARA ==========")
// 			if len(negara) == 0 {
// 				fmt.Println("Belum ada data negara nih!")
// 			} else {
// 				for kode, nama := range negara {
// 					fmt.Printf("%s: %s\n", kode, nama)
// 				}
// 			}
// 			continue
// 		}

// 		// Untuk tambah, edit, dan hapus
// 		if pilih == 1 {
// 			// Tambah negara
// 			fmt.Print("Masukkan kode negara (3 huruf): ")
// 			var kode string
// 			fmt.Scanln(&kode)
// 			kode = strings.ToUpper(kode)

// 			// Cek apakah kode sudah ada
// 			_, ada := negara[kode]
// 			if ada {
// 				fmt.Println("Kode negara udah ada!")
// 				continue
// 			}

// 			fmt.Print("Masukkan nama negara: ")
// 			var nama string
// 			fmt.Scanln(&nama)

// 			negara[kode] = nama
// 			fmt.Printf("Berhasil tambah negara %s (%s)\n", nama, kode)

// 		} else if pilih == 2 {
// 			// Edit negara
// 			fmt.Print("Masukkan kode negara yang mau diedit: ")
// 			var kode string
// 			fmt.Scanln(&kode)
// 			kode = strings.ToUpper(kode)

// 			// Cek apakah negara ada
// 			namaLama, ada := negara[kode]
// 			if !ada {
// 				fmt.Println("Negara gak ketemu!")
// 				continue
// 			}

// 			fmt.Printf("Nama lama: %s\n", namaLama)
// 			fmt.Print("Masukkan nama baru: ")
// 			var namaBaru string
// 			fmt.Scanln(&namaBaru)

// 			negara[kode] = namaBaru
// 			fmt.Printf("Berhasil update negara %s jadi %s\n", kode, namaBaru)

// 		} else if pilih == 3 {
// 			// Hapus negara
// 			fmt.Print("Masukkan kode negara yang mau dihapus: ")
// 			var kode string
// 			fmt.Scanln(&kode)
// 			kode = strings.ToUpper(kode)

// 			// Cek apakah negara ada
// 			nama, ada := negara[kode]
// 			if !ada {
// 				fmt.Println("Negara gak ketemu!")
// 				continue
// 			}

// 			delete(negara, kode)
// 			fmt.Printf("Berhasil hapus negara %s (%s)\n", nama, kode)

// 		} else {
// 			fmt.Println("Pilihan gak ada! Coba 1-5")
// 			continue
// 		}

// 		// Tampilkan data terbaru agar keliatan
// 		fmt.Println("\nData terbaru:")
// 		for kode, nama := range negara {
// 			fmt.Printf("%s: %s\n", kode, nama)
// 		}
// 	}
// }

// func main() {
// 	fmt.Println("=== PROGRAM DATA NEGARA ===")
// 	fmt.Println("Selamat datang! Yuk atur data negara")
// 	kelolaNegara()
// }
