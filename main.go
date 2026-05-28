package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.Carregar()
	router := router.Gerar()

	fmt.Println("Escutando na porta", config.Porta, "...")
	
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}