package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Rodando WebApp na porta 4000")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":4000", r))
}
