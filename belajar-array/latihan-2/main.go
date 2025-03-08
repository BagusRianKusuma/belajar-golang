package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mahasiswa struct {
	Nama    string
	Umur    int
	Jurusan string
}

func main() {
	// reader
	reader := bufio.NewReader(os.Stdin)
	var mahasiswaList []Mahasiswa // slice untuk menyimpan data

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Hapus Mahasiswa")
		fmt.Println("3. Lihat Daftar Mahasiswa")
		fmt.Println("4. Pencarian Mahasiswa")
		fmt.Println("5. Keluar")
		fmt.Println("Pilihan Menu: ")

		pilihanStr, _ := reader.ReadString('\n')
		pilihanStr = strings.TrimSpace(pilihanStr)
		pilihan, _ := strconv.Atoi(pilihanStr)

		switch pilihan {
		case 1:
			//tambahkan mahasiswa
			var mhs Mahasiswa

			fmt.Print("Masukkan Nama: ")
			mhs.Nama, _ = reader.ReadString('\n')
			mhs.Nama = strings.TrimSpace(mhs.Nama)

			fmt.Print("Masukkan Umur: ")
			umurStr, _ := reader.ReadString('\n')
			umurStr = strings.TrimSpace(umurStr)
			mhs.Umur, _ = strconv.Atoi(umurStr)

			fmt.Print("Masukkan Jurusan : ")
			mhs.Jurusan, _ = reader.ReadString('\n')
			mhs.Jurusan = strings.TrimSpace(mhs.Jurusan)

			//tambahkan ke slice mahasiswaList
			mahasiswaList = append(mahasiswaList, mhs)
			fmt.Println("Mahasiswa berhasil ditambahkan!")

		case 2:
			//hapus mahasiswa
			if len(mahasiswaList) == 0 {
				fmt.Println("Belum ada mahasiswa yang dihapus.")
				continue
			}

			fmt.Println("Masukkan nama mahasiswa yang ingin dihapus : ")
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)

			index := -1
			for i, mhs := range mahasiswaList {
				if strings.EqualFold(mhs.Nama, nama) {
					index = i
					break
				}
			}

			if index != -1 {
				mahasiswaList = append(mahasiswaList[:index], mahasiswaList[index+1:]...)
				fmt.Println("Data Berhasil Dihapus!")
			} else {
				fmt.Println("Mahasiswa Tidak Ditemukan..")
			}

		case 3:
			// Lihat Mahasiswa
			if len(mahasiswaList) == 0 {
				fmt.Println("Belum ada data mahasiswa.")
			} else {
				fmt.Println("\nDaftar Mahasiswa:")
				for i, m := range mahasiswaList {
					fmt.Printf("%d. Nama: %s | Umur: %d | Jurusan: %s\n", i+1, m.Nama, m.Umur, m.Jurusan)
				}
			}
		case 4:
			// Cari Mahasiswa
			if len(mahasiswaList) == 0 {
				fmt.Println("Belum ada data mahasiswa untuk dicari.")
				continue
			}

			fmt.Print("Masukkan Nama Mahasiswa yang ingin dicari: ")
			namaCari, _ := reader.ReadString('\n')
			namaCari = strings.TrimSpace(namaCari)

			found := false
			for _, mhs := range mahasiswaList {
				if strings.EqualFold(mhs.Nama, namaCari) {
					fmt.Printf("Ditemukan: Nama: %s | Umur: %d | Jurusan: %s\n", mhs.Nama, mhs.Umur, mhs.Jurusan)
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Mahasiswa tidak ditemukan.")
			}
		}

	}

}
