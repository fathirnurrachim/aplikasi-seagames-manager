package main

import "fmt"

func TambahMedali(data *Data, nData int) {
	var nama, jenis string
	var jumlah, index int

	fmt.Print("Nama Negara: ")
	fmt.Scan(&nama)

	index = SequentialSearch(*data, nData, nama)

	if index != -1 {

		fmt.Print("Jenis medali (emas/perak/perunggu): ")
		fmt.Scan(&jenis)

		fmt.Print("Jumlah: ")
		fmt.Scan(&jumlah)

		if jenis == "emas" {
			(*data)[index].medali.emas += jumlah
		} else if jenis == "perak" {
			(*data)[index].medali.perak += jumlah
		} else if jenis == "perunggu" {
			(*data)[index].medali.perunggu += jumlah
		}

		fmt.Println("Medali berhasil ditambah")

	} else {
		fmt.Println("Negara tidak ditemukan")
	}
}

func EditMedali(data *Data, nData int) {
	var nama, jenis string
	var jumlahBaru, index int

	fmt.Print("Nama Negara: ")
	fmt.Scan(&nama)

	index = SequentialSearch(*data, nData, nama)

	if index != -1 {

		fmt.Print("Jenis medali (emas/perak/perunggu): ")
		fmt.Scan(&jenis)

		fmt.Print("Jumlah baru: ")
		fmt.Scan(&jumlahBaru)

		if jenis == "emas" {
			(*data)[index].medali.emas = jumlahBaru
		} else if jenis == "perak" {
			(*data)[index].medali.perak = jumlahBaru
		} else if jenis == "perunggu" {
			(*data)[index].medali.perunggu = jumlahBaru
		}

		fmt.Println("Medali berhasil diedit")

	} else {
		fmt.Println("Negara tidak ditemukan")
	}
}

func HapusMedali(data *Data, nData int) {
	var nama, jenis string
	var index int

	fmt.Print("Nama Negara: ")
	fmt.Scan(&nama)

	idx = BinarySearch(*data, nData, nama)

	if index != -1 {

		fmt.Print("Jenis medali (emas/perak/perunggu): ")
		fmt.Scan(&jenis)

		if jenis == "emas" {
			(*data)[index].medali.emas = 0
		} else if jenis == "perak" {
			(*data)[index].medali.perak = 0
		} else if jenis == "perunggu" {
			(*data)[index].medali.perunggu = 0
		}

		fmt.Println("Medali berhasil dihapus")

	} else {
		fmt.Println("Negara tidak ditemukan")
	}
}
