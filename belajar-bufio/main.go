package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama Kamu : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Asal Kota : ")
	kota, _ := reader.ReadString('\n')
	kota = strings.TrimSpace(kota)

	fmt.Printf("Nama kamu adalah %s, dan asal kota kamu adalah %s", nama, kota)
}
