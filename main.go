package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	fmt.Print("Conectando a puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
