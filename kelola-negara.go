package main

import (
	"fmt"
)

func BacaNegara(data *Data, nData *int) {
	fmt.Print("Masukkan Jumlah Negara: ")
	fmt.Scan(nData)

	for i := 0; i < *nData; i++ {
		(*data)[i].id = i + 1

		fmt.Printf("Masukkan Nama Negara ke-%d: ", i+1)
		fmt.Scan(&(*data)[i].nama)
	}
}

func TambahNegara(data *Data, nData *int) {
	var jumlah int
	var namaBaru string
	var i int

	fmt.Print("Masukkan Jumlah Negara: ")
	fmt.Scan(&jumlah)

	if jumlah > NMAX-*nData {
		jumlah = NMAX - *nData
	}

	i = 0
	for i < jumlah {

		fmt.Print("Nama Negara Baru: ")
		fmt.Scan(&namaBaru)

		if SequentialSearch(*data, *nData, namaBaru) == -1 {

			(*data)[*nData].id = *nData + 1
			(*data)[*nData].nama = namaBaru

			(*data)[*nData].medali.emas = 0
			(*data)[*nData].medali.perak = 0
			(*data)[*nData].medali.perunggu = 0

			*nData = *nData + 1

			fmt.Println("NEGARA BERHASIL DITAMBAHKAN!")
		} else {
			fmt.Println("NEGARA SUDAH ADA!")
		}

		i = i + 1
	}
}

func EditNegara(data *Data, nData int) {
	var namaLama, namaBaru string
	var index int

	fmt.Print("Nama Negara Lama: ")
	fmt.Scan(&namaLama)

	index = SequentialSearch(*data, nData, namaLama)

	if index != -1 {

		fmt.Print("Nama Negara Baru: ")
		fmt.Scan(&namaBaru)

		(*data)[index].nama = namaBaru

		fmt.Println("NEGARA BERHASIL DIEDIT!")
	} else {
		fmt.Println("NEGARA TIDAK DITEMUKAN!")
	}
}

func HapusNegara(data *Data, nData *int) {
	var namaCari string
	var i, index int

	fmt.Print("Nama Negara: ")
	fmt.Scan(&namaCari)

	index = BinarySearch(*data, *nData, namaCari)

	if index != -1 {

		i = index
		for i < *nData-1 {
			(*data)[i] = (*data)[i+1]
			i = i + 1
		}

		*nData = *nData - 1

		fmt.Println("NEGARA BERHASIL DIHAPUS!")
	} else {
		fmt.Println("NEGARA TIDAK DITEMUKAN!")
	}
}

func TampilkanNegara(data Data, nData int) {
	fmt.Printf("%-5s %-25s %-10s %-10s %-10s\n",
		"ID", "Nama Negara", "Emas", "Perak", "Perunggu")

	for i := 0; i < nData; i++ {
		fmt.Printf("%-5d %-25s %-10d %-10d %-10d\n",
			data[i].id,
			data[i].nama,
			data[i].medali.emas,
			data[i].medali.perak,
			data[i].medali.perunggu)
	}
}
