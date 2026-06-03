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

	InsertionSortAsc(data, *nData)
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
