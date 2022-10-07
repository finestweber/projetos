package rota

import (
	"go-api/src/controllers"
	"net/http"
)

var RotasUsuarios = []Rota{
	{
		URI: "/home",
		Metodo: http.MethodGet,
		Funcao: controllers.Hello,
	},
}