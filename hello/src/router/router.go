package router

import (
	"go-api/src/router/rota"
	"github.com/gorilla/mux"
)

func Carrega() *mux.Router{
	r := mux.NewRouter()
	return rota.Configurar(r)
}