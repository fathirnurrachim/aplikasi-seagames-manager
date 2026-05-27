package main

import (
	"fmt"
	"strings"
)

// Data negara dan medalinya
var negara = map[string]map[string]int{
	"INA": {"emas": 0, "perak": 0, "perunggu": 0},
	"MAS": {"emas": 0, "perak": 0, "perunggu": 0},
	"SGP": {"emas": 0, "perak": 0, "perunggu": 0},
}

func kelolaMedali() {
	for {
		fmt.Println("\n========== MENU MEDALI ==========")
		fmt.Println("1. Tambah medali")
		fmt.Println("2. Edit medali")
		fmt.Println("3. Hapus medali")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih menu (1-4): ")
		
		var pilih int
		fmt.Scanln(&pilih)
		
		if pilih == 4 {
			fmt.Println("Kembali ke menu utama...")
			break
		}
		
		fmt.Print("Masukkan kode negara (INA/MAS/SGP): ")
		var kode string
		fmt.Scanln(&kode)
		kode = strings.ToUpper(kode)
		
		// Cek apakah negara ada
		_, ada := negara[kode]
		if !ada {
			fmt.Println("Negara gak ketemu!")
			continue
		}
		
		fmt.Print("Jenis medali (emas/perak/perunggu): ")
		var jenis string
		fmt.Scanln(&jenis)
		jenis = strings.ToLower(jenis)
		
		// Cek apakah jenis medali valid
		if jenis != "emas" && jenis != "perak" && jenis != "perunggu" {
			fmt.Println("Jenis medali salah! Pilih emas/perak/perunggu")
			continue
		}
		
		if pilih == 1 {
			// Tambah medali
			fmt.Print("Jumlah yang mau ditambah: ")
			var jumlah int
			fmt.Scanln(&jumlah)
			negara[kode][jenis] += jumlah
			fmt.Printf("Berhasil tambah %d %s untuk %s\n", jumlah, jenis, kode)
			
		} else if pilih == 2 {
			// Edit medali
			fmt.Print("Jumlah baru: ")
			var jumlahBaru int
			fmt.Scanln(&jumlahBaru)
			negara[kode][jenis] = jumlahBaru
			fmt.Printf("Berhasil update %s untuk %s jadi %d\n", jenis, kode, jumlahBaru)
			
		} else if pilih == 3 {
			// Hapus medali
			negara[kode][jenis] = 0
			fmt.Printf("Berhasil hapus semua %s untuk %s\n", jenis, kode)
		}
		
		// Tampilkan data terbaru biar keliatan
		fmt.Printf("\nData terbaru %s: Emas=%d, Perak=%d, Perunggu=%d\n", 
			kode, 
			negara[kode]["emas"], 
			negara[kode]["perak"], 
			negara[kode]["perunggu"])
	}
}

func main() {
	fmt.Println("=== PROGRAM DATA MEDALI OLIMPIADE ===")
	kelolaMedali()
}





