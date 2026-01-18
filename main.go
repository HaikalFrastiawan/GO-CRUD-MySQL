package main

import (
	"log"
	"net/http"
	"go-web-native/config"
	"go-web-native/controller/categorycontroller"
	"go-web-native/controller/homecontroller"
	"go-web-native/controller/productcontroller"
	
)

func main() {

	config.ConnectDB()

	//1.Home page
	http.HandleFunc("/", homecontroller.Welcome)

	//2.category page
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//product
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/delete", productcontroller.Delete)


	log.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
