package main

import (
	"fmt"
)

func main() {
	buah := []string{"apel", "mangga", "leci"}
	fmt.Println("hasil dari slice buah : ", buah)

	//menambahkan elemen baru ke slice
	buah = append(buah, "semangka", "rambutan")
	fmt.Println("ini adalah hasil dari menambahkan element dengan slice metode append ", buah)

	//hapus elemen
	index := 1
	buah = append(buah[:1], buah[index+1:]...)
	fmt.Println("ini adalah hasil dari hapus element dengan append ", buah)

	//hasil akhir
	fmt.Println("Hasil akhir dari ini adalah : ", buah)
}
