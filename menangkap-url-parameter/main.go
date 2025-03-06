package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func AgeCheck(w http.ResponseWriter, r *http.Request) {
	ageStr := r.URL.Query().Get("age")
	age, err := strconv.Atoi(ageStr)

	if err != nil {
		fmt.Fprint(w, "Silahkan Masukkan Umur yang Valid")
	} else if age > 17 {
		fmt.Fprint(w, "Anda SUdah Dewasa!")
	} else if age > 15 {
		fmt.Fprint(w, "Anda Masih Remaja")
	} else {
		fmt.Fprint(w, "anda masih anak-anak")
	}
}

func main() {
	http.HandleFunc("/age", AgeCheck)
	fmt.Println("Server Berjalan di http:localhost:8080")
	http.ListenAndServe(":8080", nil)
}
