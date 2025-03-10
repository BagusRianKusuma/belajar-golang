package main

import (
	"fmt"
)

// Fungsi untuk menjumlahkan elemen array
func penjumlahan() {
	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Ini adalah arraynya:", numbers)
	total := 0

	for _, num := range numbers {
		total += num
	}

	fmt.Println("Total:", total)
}

// Fungsi untuk mencari nilai tertinggi dan terendah dalam array
func nilaiTertinggiTerendah() {
	numbers := [6]int{3, 7, 2, 9, 5, 1}

	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	fmt.Println("Nilai Tertinggi:", max)
	fmt.Println("Nilai Terendah:", min)
}

// Fungsi untuk membalik array
func balikArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return arr
}

func main() {
	// Deklarasi array sebelum digunakan
	numbers := []int{1, 2, 3, 4, 5}

	// Memanggil fungsi balikArray
	numbers = balikArray(numbers)
	fmt.Println("Setelah Dibalik:", numbers)

	// Memanggil fungsi lainnya
	penjumlahan()
	nilaiTertinggiTerendah()
}
