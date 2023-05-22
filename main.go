package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Executando api")

	r := router.Gerar()
	log.Println("Rota iniciada")
	log.Fatal(http.ListenAndServe(":5000", r))
}
