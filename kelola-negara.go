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

	fmt.Print("Masukkan Jumlah Negara: ")
	fmt.Scan(&jumlah)

	if jumlah > NMAX-*nData {
		jumlah = NMAX - *nData
	}

	for i := 0; i < jumlah; i++ {
		fmt.Print("Nama Negara Baru: ")
		fmt.Scan(&namaBaru)

		ada := false

		for j := 0; j < *nData; j++ {
			if (*data)[j].nama == namaBaru {
				ada = true
				break
			}
		}

		if ada {
			fmt.Println("NEGARA SUDAH ADA!")
			continue
		}

		(*data)[*nData].id = *nData + 1
		(*data)[*nData].nama = namaBaru

		(*data)[*nData].medali.emas = 0
		(*data)[*nData].medali.perak = 0
		(*data)[*nData].medali.perunggu = 0

		*nData++

		fmt.Println("NEGARA BERHASIL DITAMBAHKAN!")
	}
}

func EditNegara(data *Data, nData int) {
	var namaLama, namaBaru string

	fmt.Print("Nama Negara Lama: ")
	fmt.Scan(&namaLama)

	for i := 0; i < nData; i++ {
		if (*data)[i].nama == namaLama {
			fmt.Print("Nama Negara Baru: ")
			fmt.Scan(&namaBaru)

			(*data)[i].nama = namaBaru

			fmt.Println("NEGARA BERHASIL DIEDIT!")
			return
		}
	}

	fmt.Println("NEGARA TIDAK DITEMUKAN!")
}

func HapusNegara(data *Data, nData *int) {
	var namaCari string

	fmt.Print("Nama Negara: ")
	fmt.Scan(&namaCari)

	index := -1

	for i := 0; i < *nData; i++ {
		if (*data)[i].nama == namaCari {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("NEGARA TIDAK DITEMUKAN!")
		return
	}

	for i := index; i < *nData-1; i++ {
		(*data)[i] = (*data)[i+1]
	}

	*nData--

	fmt.Println("NEGARA BERHASIL DIHAPUS!")
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
