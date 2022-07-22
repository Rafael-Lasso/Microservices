package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func loadData() []byte {

	jsonFile, err := os.Open("products.json")
	
	if err != nil {
		fmt.Println("erro: ",err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	return data
}

func ListProducts(w http.ResponseWriter, r *http.Request) {

	products := loadData()
	w.Write([]byte(products))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/products", ListProducts)
	http.ListenAndServe(":8081", r)

	fmt.Println(string(loadData()))
}