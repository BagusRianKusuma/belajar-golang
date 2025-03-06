package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json : "message"`
}

func AgeCheck(w http.ResponseWriter, r *http.Request) {
	ageStr := r.URL.Query().Get("age")
	age, err := strconv.Atoi(ageStr)

	//set response type ke JSON
	w.Header().Set("Content-Type", "application/json")

	var response Response

	if err != nil {
		response = Response{Message: "Silahkan masukkan umur yang valid"}
	} else if age > 17 {
		response = Response{Message: "Anda sudah Dewasa!"}
	} else if age > 15 {
		response = Response{Message: "Anda masih remaja!"}
	} else {
		response = Response{Message: "Anda masih anak-anak!"}
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/age", AgeCheck)
	http.ListenAndServe(":8080", nil)
}
