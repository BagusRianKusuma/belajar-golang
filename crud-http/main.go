package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// model produk
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Data Produk Dummy
var products = []Product{
	{ID: 1, Name: "Laptop", Price: 5000000},
	{ID: 2, Name: "Keyboard", Price: 1500000},
	{ID: 3, Name: "Mouse", Price: 750000},
}

// helper untuk mengirim response JSON
func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// read(GET/Products) - Ambil semua produk
func getProducts(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, http.StatusOK, products)
}

// create (POST/Products) - Tambah Produk Baru
func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request paylod", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(products) + 1 // auto increment ID
	products = append(products, newProduct)

	sendJSON(w, http.StatusCreated, newProduct)
}

// update(PUT/products/{id})update produk berdasarkan ID
func updateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	for i, p := range products {
		if p.ID == id {
			products[i].Name = updatedProduct.Name
			products[i].Price = updatedProduct.Price
			sendJSON(w, http.StatusOK, products[i])
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}

// Delete (DELETE /products/{id}) - Hapus produk
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...) // hapus produk
			sendJSON(w, http.StatusOK, map[string]string{"message": "Product Deleted"})
			return
		}
	}

	http.Error(w, "Product Not Found!", http.StatusNotFound)
}

// routes di main
func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProducts(w, r)
		case http.MethodPost:
			createProduct(w, r)
		case http.MethodPut:
			updateProduct(w, r)
		case http.MethodDelete:
			deleteProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
