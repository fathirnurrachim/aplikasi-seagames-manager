package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Negara struct {
	Nama string
	Kode string
}

var daftarNegara []Negara
var sc = bufio.NewScanner(os.Stdin)

func kelolaNegara() {
	for {
		fmt.Println("\n1. Tambah Negara")
		fmt.Println("2. Edit Negara")
		fmt.Println("3. Hapus Negara")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih: ")

		var pilih int
		fmt.Scanln(&pilih)
		sc.Scan()

		if pilih == 4 {
			break
		}

		switch pilih {
		case 1:
			fmt.Print("Nama Negara: ")
			nama := bacaString()
			fmt.Print("Kode Negara: ")
			kode := bacaString()

			ada := false
			for _, n := range daftarNegara {
				if n.Kode == kode {
					ada = true
					break
				}
			}

			if !ada {
				daftarNegara = append(daftarNegara, Negara{nama, kode})
				fmt.Println("Berhasil ditambahkan")
			} else {
				fmt.Println("Kode sudah ada")
			}

		case 2:
			fmt.Print("Kode negara yang diedit: ")
			kode := bacaString()

			ketemu := -1
			for i, n := range daftarNegara {
				if n.Kode == kode {
					ketemu = i
					break
				}
			}

			if ketemu != -1 {
				fmt.Print("Nama baru: ")
				daftarNegara[ketemu].Nama = bacaString()
				fmt.Println("Berhasil diupdate")
			} else {
				fmt.Println("Negara tidak ditemukan")
			}

		case 3:
			fmt.Print("Kode negara yang dihapus: ")
			kode := bacaString()

			ketemu := -1
			for i, n := range daftarNegara {
				if n.Kode == kode {
					ketemu = i
					break
				}
			}

			if ketemu != -1 {
				daftarNegara = append(daftarNegara[:ketemu], daftarNegara[ketemu+1:]...)
				fmt.Println("Berhasil dihapus")
			} else {
				fmt.Println("Negara tidak ditemukan")
			}

		default:
			fmt.Println("Pilihan salah")
		}
	}
}

func bacaString() string {
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}

func main() {
	kelolaNegara()
}
