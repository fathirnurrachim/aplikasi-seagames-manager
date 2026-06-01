package main

import (
	"fmt"
	// "strings"
)

func BacaDataNegara(data *Data, nData *int) {

	fmt.Print("Masukkan Banyaknya Negara: ")
	fmt.Scan(&nData)

	if *nData > NMAX {
		*nData = NMAX
	}

	for i := 0; i < *nData; i++ {
		fmt.Print("Masukkan ID Negara: ")
		fmt.Scan(&data[i].id)

		fmt.Print("Masukkan Nama Negara: ")
		fmt.Scan(&data[i].nama)
	}
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
