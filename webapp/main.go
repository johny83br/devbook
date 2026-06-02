package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Rodando WebApp na porta 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
