package config

import (
	"github.com/joho/godotenv"
	"log"
	"strconv"
	"os"
	
)

var Porta = 0

func Carregar(){
	var erro error
	if erro = godotenv.Load(); erro != nil{
		log.Fatal(erro)
	}
	if Porta, erro = strconv.Atoi(os.Getenv("API_PORT")); erro != nil{
		Porta = 1234
	}
}