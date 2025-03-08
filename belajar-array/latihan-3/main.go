/*
Ini adalah latihan slice dengan fmt.scan
*/

package main

import (
	"fmt"
	"strings"
)

type Buah struct {
	Nama string
}

func tambahBuah(buah []Buah, nama string) []Buah {
	buah = append(buah, Buah{Nama: nama})
	return buah
}

func hapusBuah(buah []Buah, nama string) []Buah {
	index := -1
	for i, b := range buah {
		if strings.EqualFold(b.Nama, nama) {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Buah tidak ditemukan!")
		return buah
	}

	return append(buah[:index], buah[index+1:]...)
}

func main() {
	//slice awal
	buahList := []Buah{
		{"apel"},
		{"kelengkeng"},
		{"semangka"},
	}

	//menambah buah
	var buahBaru string
	fmt.Print("Masukan Nama Buah yang ingin ditambahkan: ")
	fmt.Scan(&buahBaru)
	buahList = tambahBuah(buahList, buahBaru)

	fmt.Println("Daftar buah setelah ditambah : ", buahList)

	//menghapus buah
	var buahHapus string
	fmt.Print("Masukkan Nama buah yang ingin dihapus : ")
	fmt.Scan(&buahHapus)
	buahList = hapusBuah(buahList, buahHapus)

	fmt.Println("Daftar buah setelah dihapus : ", buahList)
}
