package main

import (
	"fmt"
)

<<<<<<< HEAD
func BacaNegara(data *Data, nData *int) {
	fmt.Print("Masukkan Jumlah Negara: ")
	fmt.Scan(nData)

	for i := 0; i < *nData; i++ {
		data[i].id = generateID(*data, i)
		fmt.Printf("Masukkan Nama Negara ke-%d: ", i+1)
		fmt.Scan(&data[i].nama)
=======
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
>>>>>>> 892f627defad48df4db8faa02736c861d94b4e2f
	}
}

func TampilkanNegara(data Data, nData int) {

	fmt.Printf("%-5s %-25s %-10s %-10s %-10s\n",
		"ID", "Nama Negara", "Emas", "Perak", "Perunggu")

	for i := 0; i < nData; i++ {
		fmt.Printf("%-5d %-25s %-10d %-10d %-10d\n",
			data[i].id, data[i].nama, data[i].medali.emas, data[i].medali.perak, data[i].medali.perunggu)
	}
}

func TambahNegara(data *Data, nData *int, idBaru int, namaBaru string) {
	var index int

	index = SequentialSearch(*data, *nData, idBaru)

	if index == -1 {
		(*data)[*nData].id = idBaru
		(*data)[*nData].nama = namaBaru

		(*data)[*nData].medali.emas = 0
		(*data)[*nData].medali.perak = 0
		(*data)[*nData].medali.perunggu = 0

		*nData++

		fmt.Println("Negara Berhasil Ditambahkan!")
	} else {
		fmt.Println("Negara Sudah Ada!")
	}
}

func EditNegara(data *Data, nData int, idCari int, namaBaru string) {
	var index int

	index = SequentialSearch(*data, nData, idCari)

	if index != -1 {
		data[index].nama = namaBaru
		fmt.Println("Negara Berhasil Diedit!")
	} else {
		fmt.Println("Negara yang Ingin Diedit Tidak Ditemukan!")
	}
}

func HapusNegara(data *Data, nData *int, idCari int) {
	var index int

	index = BinarySearch(*data, *nData, idCari)

	if index != -1 {
		for i := index; i < *nData-1; i++ {
			data[i] = data[i+1]
			fmt.Println("Negara Berhasil Dihapus!")
		}
		*nData--
	} else {
		fmt.Println("Negara yang Ingin Dihapus Tidak Ditemukan!")
	}
}
