package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	Metodo:  http.MethodPost,
	URI:    "/login",
	Funcao: controllers.Login,
	RequerAutenticacao: false,
}