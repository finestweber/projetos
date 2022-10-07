
// variaveis de ambiente
package main

import (
	"go-api/src/config"
	"go-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Println("Servidor Rodando na porta 5000")
	config.Carregar()
	
	r := router.Carrega()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}