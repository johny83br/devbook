package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Rodando WebApp na porta", config.APP_PORTA)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.APP_PORTA), r))
}
