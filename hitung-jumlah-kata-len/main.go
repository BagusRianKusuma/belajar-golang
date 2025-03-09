package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan sebuah kalimat : ")
	kalimat, _ := reader.ReadString('\n')
	kalimat = strings.TrimSpace(kalimat)

	kataList := strings.Fields(kalimat)
	jumlahKata := len(kataList)

	fmt.Println("Jumlah kata dalam kalimat anda adalah : ", jumlahKata)
}
