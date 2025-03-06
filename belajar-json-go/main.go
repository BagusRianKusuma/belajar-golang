package main

import (
	"encoding/json"
	"net/http"
)

// struct untuk response JSON
type Response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func sendJSON(w http.ResponseWriter, r *http.Request) {
	//Data yang akan dikirim dalam bentuk JSON
	data := Response{
		Name: "Bagus Rian Kusuma",
		Age:  24,
	}

	// set response header agar formatnya JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode struct menjadi JSON dan kirim ke client
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/json", sendJSON)
	http.ListenAndServe(":8080", nil)
}
