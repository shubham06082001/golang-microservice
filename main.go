package main

import (
	"log"
	"net/http"
	"os"
	"C:\Users\e019783\Desktop\microservice\handlers"
)

func main(){

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.newHello(l)

	http.HandleFunc()

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request){
		log.Println("Hello World!!")
	})

	http.ListenAndServe(":9090", nil)
}