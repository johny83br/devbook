package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	API_URL   = "http://localhost"
	API_PORTA = 5000
	APP_PORTA = 4000
	HASH_KEY  []byte
	BLOCK_KEY []byte
)

func Carregar() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	API_PORTA, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		log.Fatal("Erro ao converter a porta para inteiro:", erro)
	}

	APP_PORTA, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal("Erro ao converter a porta para inteiro:", erro)
	}

	API_URL = os.Getenv("API_URL")
	HASH_KEY = []byte(os.Getenv("HASH_KEY"))
	BLOCK_KEY = []byte(os.Getenv("BLOCK_KEY"))

}
